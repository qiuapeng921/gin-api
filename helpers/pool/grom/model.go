package grom

import (
	"fmt"
	"gin-api/app/models/user"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

var client *gorm.DB

func SetUpOrm() {
	database := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_CHARSET"))
	var err error
	client, err = gorm.Open("mysql", database)
	if err != nil {
		panic(err.Error())
	}
	client.SingularTable(true)
	maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	client.DB().SetMaxIdleConns(maxIdle)
	client.DB().SetMaxOpenConns(maxOpen)
	client.LogMode(true)
	fmt.Println("mysql连接成功")

	client.AutoMigrate(&user.Users{})
}

func GetOrm() *gorm.DB {
	return client
}