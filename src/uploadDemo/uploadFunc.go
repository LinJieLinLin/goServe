package uploadDemo

import (
	"mime/multipart"
	"strings"
	"config"
	"time"
	"io"
	"net/http"
	"github.com/Centny/Cny4go/log"
	"fmt"
	"os"
	"errors"
	"strconv"
)

type ReData struct {
	Msg     string `json:"msg"`
	Code    int64 `json:"code"`
	Data    interface{} `json:"data"`
}

// 获取文件大小的接口
type Size interface {
	Size() int64
}

// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

// hello world, the web server
const (
	picType = `jpg,png,jpeg,bmp` //可以上传的文件格式
)

func CheckFile(file multipart.File, fileHeader *multipart.FileHeader, r *http.Request) (string, error) {
	log.I("start CheckFile")
	//检查文件格式
	name := strings.Split(fileHeader.Filename, ".")
	if !strings.Contains(picType, name[len(name)-1]) {
		log.E("文件格式不符合")
		return "file type is  wrong",errors.New("文件格式不符合")
	}
	//检查文件大小   末完成
	if statInterface, ok := file.(Size); ok {
		sizeB:= statInterface.Size()
		//保留小数点后两位
		sizeM, _ := strconv.ParseFloat(strconv.FormatFloat(float64(sizeB) / (1024*1024), 'f', 2, 64), 64)
		fmt.Println(sizeB,"B", sizeM, "M")
		log.I("获取上传的文件大小")
	}
	//读取上传路径
	var PRE_IMG_SAVE_PATH string = config.GetConfig("FILE_PATH").(string)
	//更改文件名
	fileName := fmt.Sprintf("%d", time.Now().UnixNano()) + fileHeader.Filename
	filePath := PRE_IMG_SAVE_PATH + "picFile/"
	imgPaths := filePath + fileName
	//创建目录
	err := makeDir(filePath)
	if err != nil {
		return imgPaths, err
	}
	if err := saveFile(file, imgPaths); err != nil {
		log.E("保存文件出错：", err, imgPaths)
		return "save File picture errror", err
	}
	//去除www
	imgPaths = strings.Replace(imgPaths, "www", "", -1)
	return imgPaths, nil
}
/*
	  递归创建目录
	  os.MkdirAll(path string, perm FileMode) error

	  path  目录名及子目录
	  perm  目录权限位
	  error 如果成功返回nil，如果目录已经存在默认什么都不做
	*/
func makeDir(arg_dir string) (error) {
	dirPath := arg_dir
	if err := os.MkdirAll(dirPath,0); err!= nil {
		log.E("创建目录失败", err)
		return err
	}
	return nil
}
//保存文件
func saveFile(arg_File multipart.File, arg_filePath string) error {
	defer arg_File.Close()
	if f, err := os.OpenFile(arg_filePath, os.O_WRONLY|os.O_CREATE, 0666); err != nil {
		log.E("打开文件失败",err)
		return err
	} else {
		if _, err := io.Copy(f, arg_File); err == nil {
			f.Close()
		} else {
			log.E("复制文件失败",err)
			return err
		}
	}
	return nil
}
