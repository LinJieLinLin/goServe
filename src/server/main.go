package server

import (
	"config"
	"databaseConn"
	"github.com/gorilla/mux"//mux.NewRouter()
	"log"
	"net/http"
	"fmt"
	cny4goLog"github.com/Centny/Cny4go/log"
	"os"
	"bufio"
	"github.com/Centny/Cny4go/util"
	"io"
	"sync"
	"test"
	"uploadDemo"
)

var onceApp sync.Once

func routeConfig() {
//
	//这里写初始运行的函数
//
	//open a timeKeeper to chk untreated message

	log.SetFlags(log.Llongfile)
	r := mux.NewRouter()
//	r.HandleFunc("/index/getAdvertisementPic",firstPage.GetAdvertisementPic).Methods("POST")
//	r.HandleFunc("/index/getItemPic", firstPage.GetFirstPagePic).Methods("POST")
	r.HandleFunc("/test",test.Test)
	r.HandleFunc("/uploadDemo",uploadDemo.Upload)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("www/")))
	http.Handle("/", r)
}

func startOnce() {
	go func() {
		//端口
		srvPort := config.GetConfig("SERVER_PORT")
		//数据库连接
		dbConfig := config.GetConfig("DB_CONFIG")
//日志部分
		logPath := config.GetConfig("LOG_PATH")
		logName := config.GetConfig("LOG_FILE_NAME")
//		create and open log file
		fmt.Println("log path : "+ logPath.(string)+logName.(string) )
		util.FTouch(logPath.(string)+logName.(string))
		f, err := os.OpenFile(logPath.(string)+logName.(string), os.O_RDWR  | os.O_CREATE  |os.O_APPEND, 0666)
		if err != nil {
			panic(err.Error())
			return
		}
		//create file buffer writer.
		bo := bufio.NewWriter(f)
		defer f.Close()
		defer bo.Flush()
		cny4goLog.SetWriter(io.MultiWriter(bo, os.Stdout))
//日志部分
		databaseConn.SetConnConfig(dbConfig.(string))
		databaseConn.GetNewConn()

		routeConfig()
		if err := http.ListenAndServe(":" + srvPort.(string), nil); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}

func startTestOnce() {
	go func() {
		routeConfig()
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}

func Start() {
	fmt.Println("成功启动！")
	onceApp.Do(startOnce)
}

func StartTestServer() {
	onceApp.Do(startTestOnce)
}
