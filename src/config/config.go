package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"strings"
	"net/http"
	"github.com/Centny/Cny4go/log"
)

var configMap map[string]interface{} = make(map[string]interface{})
var configLock *sync.RWMutex = new(sync.RWMutex)

func ServerConfig() {
	// pwd, _ := os.Getwd()
	// exec.Command("export", "GOPATH="+pwd+":$GOPATH")
	// gopath := os.Getenv("GOPATH")
	// project_dir := strings.Split(gopath, ":")[0]
	project_dir, _ := os.Getwd()
	fmt.Println("工程目录",project_dir)
	project_dir = strings.Replace(strings.SplitAfterN(project_dir, "src", -1)[0], "/src", "", 1)
	project_dir = strings.Replace(project_dir, "\\src", "", 1)
	fmt.Println("project dir is ", project_dir)
	configMap["PROJECT_DIR"] = project_dir
	if config_byte, err := ioutil.ReadFile(project_dir + "/config/config.json"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(json.Unmarshal(config_byte, &configMap))
	}
	fmt.Println("server config:")
	for k, v := range configMap {
		fmt.Printf("%s:%v \n", k, v)
	}

}

func GetConfig(k string) interface{} {
	configLock.RLock()
	defer configLock.RUnlock()
	v := configMap[k]
	return v
}

func SetConfig(k string, v interface{}) {
	configLock.Lock()
	defer configLock.Unlock()
	configMap[k] = v
}
//返回数据
func Send(v interface {}, w http.ResponseWriter) {
	b, err := json.Marshal(v)
	if err != nil {
		log.E("返回数据有误！",err)
		return
	}
	fmt.Fprintf(w, string(b))
}



func init() {
	ServerConfig()
}
