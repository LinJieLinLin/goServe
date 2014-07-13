package util

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//Author:Centny
//
//the file configure
//
type Fcfg map[string]interface{}

//get the value by key.
func (f *Fcfg) Val(key string) string {
	if val, ok := (*f)[key]; ok {
		return val.(string)
	} else {
		return ""
	}
}

//get the int value by key.
func (f *Fcfg) IntVal(key string) int {
	if !f.Exist(key) {
		return 0
	}
	val, err := strconv.Atoi(f.Val(key))
	if err != nil {
		panic(err)
	}
	return val
}

//get the float value by key.
func (f *Fcfg) FloatVal(key string) float64 {
	if !f.Exist(key) {
		return 0
	}
	val, err := strconv.ParseFloat(f.Val(key), 64)
	if err != nil {
		panic(err)
	}
	return val
}

//set the value by key and value.
func (f *Fcfg) SetVal(key string, val string) *Fcfg {
	(*f)[key] = val
	return f
}

//delete the value by key.
func (f *Fcfg) Del(key string) *Fcfg {
	delete(*f, key)
	return f
}

//check if exist by key.
func (f *Fcfg) Exist(key string) bool {
	_, ok := (*f)[key]
	return ok
}

//initial the configure by .properties file.
func (f *Fcfg) InitWithFilePath(fp string) error {
	if !Fexists(fp) {
		return errors.New("file not found")
	}
	fh, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer fh.Close()
	return f.InitWithFile(fh)
}

//initial the configure by .properties file.
func (f *Fcfg) InitWithFile(tfile *os.File) error {
	reader := bufio.NewReader(tfile)
	buf := bytes.NewBuffer(make([]byte, 0, 10240))
	for {
		//read one line
		buf.Reset()
		for {
			pair, isPrefix, err := reader.ReadLine()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			buf.Write(pair)
			if !isPrefix {
				break
			}
		}
		//
		line := buf.String()
		line = strings.Trim(line, " ")
		if len(line) < 1 {
			continue
		}
		ps := strings.Split(line, "#")
		if len(ps) < 1 || len(ps[0]) < 1 {
			continue
		}
		line = ps[0]
		ps = strings.Split(line, "=")
		if len(ps) < 2 {
			continue
		}
		key := f.EnvReplace(strings.Trim(ps[0], " "))
		val := f.EnvReplace(strings.Trim(ps[1], " "))
		(*f)[key] = val
	}
	return nil
}

//replace tartget patter by ${key} with value in configure map or system environment value.
func (f *Fcfg) EnvReplace(val string) string {
	reg, err := regexp.Compile("\\$\\{[^\\}]*\\}")
	if err != nil {
		panic(err)
	}
	var rval string = ""
	mhs := reg.FindAll([]byte(val), -1)
	for i := 0; i < len(mhs); i++ {
		bys := mhs[i]
		ptn := string(bys)
		bys = bys[2 : len(bys)-1]
		if len(bys) < 1 {
			continue
		}
		key := string(bys)
		if f.Exist(key) {
			rval = f.Val(key)
		} else {
			rval = os.Getenv(key)
		}
		if len(rval) < 1 {
			continue
		}
		val = strings.Replace(val, ptn, rval, 1)
	}
	return val
}
