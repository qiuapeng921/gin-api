package admin

import (
	"fmt"
	"gin-api/app/models/admin_role"
	"gin-api/app/models/admins"
	"gin-api/app/utility/db"
	"gin-api/app/utility/response"
	"gin-api/app/utility/system"
	"github.com/gin-gonic/gin"
	"strconv"
)

type adminRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Status   int    `json:"status"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	RoleId   int    `json:"role_id" form:"role_id" binding:"required"`
}

func Detail(ctx *gin.Context) {
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
	var request adminRequest
	var err error
	if err = ctx.ShouldBind(&request); err != nil {
		response.Context(ctx).Error(10000, "参数验证错误:"+err.Error())
		return
	}
	var admin admins.Entity

	// 判断用户是否存在
	if adminCount, _ := db.OrmClient().Where("username=?", request.Username).Count(&admin); adminCount > 0 {
		response.Context(ctx).Error(10001, "用户"+request.Username+"已存在")
		return
	}

	admin.Username = request.Username
	admin.Password = system.EncodeMD5(request.Username)
	admin.Phone = request.Phone

	// 开启一个事物管道
	session := db.OrmClient().NewSession()
	defer session.Close()

	// 开启事务
	if err := session.Begin(); err != nil {
		response.Context(ctx).Error(10002, err.Error())
		return
	}

	// 添加管理员
	_, err = db.OrmClient().Insert(&admin)
	if err != nil {
		session.Rollback()
		response.Context(ctx).Error(10003, "添加管理员失败:"+err.Error())
		return
	}

	var role admin_role.Entity
	role.AdminId = admin.Id
	role.RoleId = request.RoleId

	// 添加管理员角色关系
	_, err = session.Insert(&role)
	if err != nil {
		session.Rollback()
		response.Context(ctx).Error(10005, "管理员角色添加失败:"+err.Error())
		return
	}
	session.Commit()
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
		request adminRequest
		err     error
		admin   admins.Entity
	)

	// 判断用户是否存在
	_, _ = db.OrmClient().Where("id=?", id).Get(&admin)
	if admin.Id == 0 {
		response.Context(ctx).Error(10001, "用户"+queryId+"不存在")
		return
	}

	if err = ctx.ShouldBind(&request); err != nil {
		response.Context(ctx).Error(10002, "参数验证错误:"+err.Error())
		return
	}

	if request.Username != admin.Username {
		fmt.Println(request, admin)
		_, err := db.OrmClient().Where("id=?", id).Count(&admin)
		if err != nil {
			response.Context(ctx).Error(10003, err.Error())
			return
		}
		fmt.Println(id, admin)
		if id != admin.Id {
			response.Context(ctx).Error(10004, request.Username+":已存在,请更换用户")
			return
		}
	}

	session := db.OrmClient().NewSession()
	defer session.Close()

	// 开启事务
	if err := session.Begin(); err != nil {
		response.Context(ctx).Error(10000, err.Error())
		return
	}

	admin.Username = request.Username
	admin.Password = system.EncodeMD5(request.Password)
	admin.Phone = request.Phone
	admin.Status = request.Status

	_, err = session.ID(id).Update(&admin)
	if err != nil {
		session.Rollback()
		response.Context(ctx).Error(10005, "更新管理员失败:"+err.Error())
		return
	}

	var adminRole admin_role.Entity
	adminRole.RoleId = request.RoleId
	if _, err = session.Where("admin_id=?", id).Update(&adminRole); err != nil {
		session.Rollback()
		response.Context(ctx).Error(10006, "更新管理员角色关系失败:"+err.Error())
		return
	}
	session.Commit()
	response.Context(ctx).Success(request)
	return
}

func DeleteAdmin() {

}
