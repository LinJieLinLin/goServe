package databaseConn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"config"
	"github.com/Centny/Cny4go/log"
)

var db  *sql.DB
var testdb  *sql.DB
var connConfig string = ""

//设置数据库连接参数
func SetConnConfig(s string) {
	connConfig = s
}
//连接数据库
func GetConn() (*sql.DB, error) {
	if db == nil {
		var err error
		db, err = sql.Open("mysql", connConfig)
		if err != nil {
			log.E("GetConn db error:", err.Error())
			return nil, err
		}
	}
	return db, nil
}
//获取新连接
func GetNewConn() (*sql.DB, error) {
	if db!= nil {
		log.I("Close the old db open the new one : %v \n",connConfig)
		db.Close()
	}
	var err error
	db, err = sql.Open("mysql", connConfig)
	if err != nil {
		log.E("GetNewConn db error:", err.Error())
		return nil, err
	}
	return db, nil
}
//关数据库
func CloseDB() {
	if db != nil {
		err := db.Close()
		db = nil
		if err != nil {
			log.E("close db error:", err.Error())
		}
	}
}

//连接测试数据库
func GetTestDB() (*sql.DB, error) {
	if testdb == nil {
		var err error
		testdb, err = sql.Open("mysql", "cny:123@tcp(192.168.1.17:3306)/rcp_test?charset=utf8")
		if err != nil {
			log.E("GetTestDB db error:", err.Error())
			return nil, err
		}
	}
	return testdb, nil
}
//重新连接测试数据库
func NewTestConn() (*sql.DB, error) {
	var conn *sql.DB
	dbConfig := config.GetConfig("TEST_DB_CONFIG")
	fmt.Println("your test db is: ")
	SetConnConfig(dbConfig.(string))
	fmt.Println(connConfig)
	conn , _ = GetNewConn()
	return conn, nil
}
//关测试数据库
func CloseTestDB() {
	if testdb != nil {
		err := testdb.Close()
		testdb = nil
		if err != nil {
			log.E("close db error:", err.Error())
		}
	}
}
