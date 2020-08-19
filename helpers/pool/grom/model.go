package grom

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var client *xorm.Engine

func SetUpOrm() {
	database := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CHARSET"))
	var err error
	client, err = xorm.NewEngine("mysql", database)
	if err != nil {
		panic(err.Error())
	}
	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	client.SetMaxIdleConns(maxIdle)
	client.SetMaxOpenConns(maxOpen)
	fmt.Println("mysql连接成功")

	client.ShowSQL(true)
	client.Logger().SetLevel(log.LOG_DEBUG)
}

func GetOrm() *xorm.Engine {
	return client
}
