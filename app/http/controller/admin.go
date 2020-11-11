package controller

import (
	"fmt"
	"gin-admin/app/http/request"
	"gin-admin/app/models/admin_role"
	"gin-admin/app/models/admins"
	"gin-admin/app/utility/app"
	"gin-admin/app/utility/response"
	"gin-admin/app/utility/system"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @获取列表页面
// @Author 邱阿朋
// @Date 12:29 2020/11/05
func GetPage(ctx *gin.Context) {
	response.Context(ctx).View("admin/list")
	return
}

func GetAdminList(ctx *gin.Context) {
	var admin admins.Entity
	_, _ = app.DB().Get(&admin)
	count, _ := app.DB().Count(&admin)
	response.Context(ctx).Page(int(count),admin)
	return
}

func GetAdminDetail(ctx *gin.Context) {
	id := ctx.GetInt("id")
	result, resErr := admins.GetAdminById(id)
	if resErr != nil {
		response.Context(ctx).Error(10001, resErr.Error())
		return
	}
	response.Context(ctx).Success(result)
	return
}

func CreateAdmin(ctx *gin.Context) {
	var params request.AdminRequest
	var err error
	if err = ctx.ShouldBind(&params); err != nil {
		response.Context(ctx).Error(10000, "参数验证错误:"+err.Error())
		return
	}
	var admin admins.Entity

	db := app.DB()
	// 判断用户是否存在
	if adminCount, _ := db.Where("username=?", params.Username).Count(&admin); adminCount > 0 {
		response.Context(ctx).Error(10001, "用户"+params.Username+"已存在")
		return
	}

	admin.Username = params.Username
	admin.Password = system.EncodeMD5(params.Username)
	admin.Phone = params.Phone

	// 开启一个事物管道
	session := db.NewSession()
	defer func() {
		_ = session.Close()
	}()

	// 开启事务
	if err := session.Begin(); err != nil {
		response.Context(ctx).Error(10002, err.Error())
		return
	}

	// 添加管理员
	_, err = session.Insert(&admin)
	if err != nil {
		_ = session.Rollback()
		response.Context(ctx).Error(10003, "添加管理员失败:"+err.Error())
		return
	}

	var role admin_role.Entity
	role.AdminId = admin.Id
	role.RoleId = params.RoleId

	// 添加管理员角色关系
	_, err = session.Insert(&role)
	if err != nil {
		_ = session.Rollback()
		response.Context(ctx).Error(10005, "管理员角色添加失败:"+err.Error())
		return
	}
	_ = session.Commit()
	response.Context(ctx).Success(gin.H{"admin": admin, "admin_role": role})
	return
}

func UpdateAdmin(ctx *gin.Context) {
	queryId := ctx.Query("id")
	if queryId == "" {
		response.Context(ctx).Error(10000, "id不能为空")
		return
	}
	id, _ := strconv.Atoi(queryId)

	var (
		params request.AdminRequest
		err    error
		admin  admins.Entity
	)

	// 判断用户是否存在
	_, _ = app.DB().Where("id=?", id).Get(&admin)
	if admin.Id == 0 {
		response.Context(ctx).Error(10001, "用户"+queryId+"不存在")
		return
	}

	if err = ctx.ShouldBind(&params); err != nil {
		response.Context(ctx).Error(10002, "参数验证错误:"+err.Error())
		return
	}

	db := app.DB()
	if params.Username != admin.Username {
		fmt.Println(params, admin)
		_, err := db.Where("id=?", id).Count(&admin)
		if err != nil {
			response.Context(ctx).Error(10003, err.Error())
			return
		}
		fmt.Println(id, admin)
		if id != admin.Id {
			response.Context(ctx).Error(10004, params.Username+":已存在,请更换用户")
			return
		}
	}

	session := db.NewSession()
	defer func() {
		_ = session.Close()
	}()

	// 开启事务
	if err := session.Begin(); err != nil {
		response.Context(ctx).Error(10000, err.Error())
		return
	}

	admin.Username = params.Username
	admin.Password = system.EncodeMD5(params.Password)
	admin.Phone = params.Phone
	admin.Status = params.Status

	_, err = session.ID(id).Update(&admin)
	if err != nil {
		_ = session.Rollback()
		response.Context(ctx).Error(10005, "更新管理员失败:"+err.Error())
		return
	}

	var adminRole admin_role.Entity
	adminRole.RoleId = params.RoleId
	if _, err = session.Where("admin_id=?", id).Update(&adminRole); err != nil {
		_ = session.Rollback()
		response.Context(ctx).Error(10006, "更新管理员角色关系失败:"+err.Error())
		return
	}
	_ = session.Commit()
	response.Context(ctx).Success(params)
	return
}

func DeleteAdmin(ctx *gin.Context) {
	response.Context(ctx).Success()
	return
}
