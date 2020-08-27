package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var mysqlClient *xorm.Engine

func SetUpXorm() {
	database := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CHARSET"))
	var err error
	mysqlClient, err = xorm.NewEngine("mysql", database)
	if err != nil {
		panic(err.Error())
	}
	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	mysqlClient.SetMaxIdleConns(maxIdle)
	mysqlClient.SetMaxOpenConns(maxOpen)
	fmt.Println("mysql连接成功")

	mysqlClient.ShowSQL(false)
	mysqlClient.Logger().SetLevel(log.LOG_DEBUG)

}

func Xorm() *xorm.Engine {
	return mysqlClient
}