package grom

import (
	"fmt"
	"gin-api/app/models/admin_role"
	"gin-api/app/models/admins"
	"gin-api/app/models/article_tags"
	"gin-api/app/models/articles"
	"gin-api/app/models/categorys"
	"gin-api/app/models/comments"
	"gin-api/app/models/menus"
	"gin-api/app/models/permissions"
	"gin-api/app/models/role_permission"
	"gin-api/app/models/roles"
	"gin-api/app/models/system_config"
	"gin-api/app/models/tags"
	"gin-api/app/models/user_behavior"
	"gin-api/app/models/users"
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

	autoGenModel()
}

func GetOrm() *xorm.Engine {
	return client
}

func autoGenModel() {
	_ = client.Sync2(
		&admins.Entity{},
		&admin_role.Entity{},
		&articles.Entity{},
		&article_tags.Entity{},
		&categorys.Entity{},
		&comments.Entity{},
		&menus.Entity{},
		&permissions.Entity{},
		&roles.Entity{},
		&role_permission.Entity{},
		&system_config.Entity{},
		&tags.Entity{},
		&users.Entity{},
		&user_behavior.Entity{},
	)
}
