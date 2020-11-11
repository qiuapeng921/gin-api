package database

import (
	"fmt"
	"gin-api/app/models/admins"
	"gin-api/app/models/users"
	"gin-api/app/utility/app"
	"gin-api/app/utility/system"
	"xorm.io/xorm"
)

var orm *xorm.Engine

func AutoGenTable() {
	orm = app.DB()
	_ = orm.Sync2(
		&admins.Entity{},
		&users.Entity{},
	)

	if result, _ := orm.IsTableEmpty(&admins.Entity{}); result {
		defaultData()
	}
}

func defaultData() {
	var admin admins.Entity
	admin.Username = "admin"
	admin.Password = system.EncodeMD5("123456")
	admin.Phone = "15249279779"
	if _, err := orm.InsertOne(&admin); err != nil {
		fmt.Println("初始化超管失败," + err.Error())
	}
	fmt.Println("初始化超管成功")
}
