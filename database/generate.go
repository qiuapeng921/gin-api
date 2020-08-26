package database

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
	"gin-api/helpers/db"
	"gin-api/helpers/system"
	"xorm.io/xorm"
)

var orm *xorm.Engine

func AutoGenTable() {
	orm = db.Xorm()
	_ = orm.Sync2(
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

	if result, _ := db.Xorm().IsTableEmpty(&admins.Entity{}); result {
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

	var role roles.Entity
	role.RoleName = "超管"
	role.RoleDesc = "超级管理员"
	if _, err := orm.Insert(&role); err != nil {
		fmt.Println("初始化角色失败," + err.Error())
	}
	fmt.Println("初始化角色成功")

	var adminRole admin_role.Entity
	adminRole.AdminId = 1
	adminRole.RoleId = 1
	if _, err := orm.Insert(&adminRole); err != nil {
		fmt.Println("初始化用户角色关系失败," + err.Error())
	}
	fmt.Println("初始化用户角色关系成功")
}
