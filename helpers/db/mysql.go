package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var mysqlClient *xorm.Engine

func InitXorm() {
	database := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CHARSET"))
	fmt.Println(database)
	var err error
	mysqlClient, err = xorm.NewEngine("mysql", database)
	if err != nil {
		panic(fmt.Sprintf("初始化数据库失败 错误信息:%s", err.Error()))
	}
	if err := mysqlClient.Ping(); err != nil {
		panic(fmt.Sprintf("连接数据库失败 错误信息:%s", err.Error()))
	}
	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	mysqlClient.SetMaxIdleConns(maxIdle)
	mysqlClient.SetMaxOpenConns(maxOpen)
	fmt.Println("mysql连接成功")

	show := os.Getenv("APP_ENV")
	if show == gin.DebugMode {
		mysqlClient.ShowSQL(true)
		mysqlClient.Logger().SetLevel(log.LOG_DEBUG)
	}
}

func OrmClient() *xorm.Engine {
	return mysqlClient
}