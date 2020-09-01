package admin

import (
	"gin-api/app/models/roles"
	"gin-api/helpers/db"
	"gin-api/helpers/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type roleRequestData struct {
	RoleName string `json:"role_name" from:"role_name" binding:"required"`
	RoleDesc string `json:"role_desc" from:"role_desc" binding:"required"`
	Status   int    `json:"status" from:"status"`
}

func RoleList(ctx *gin.Context) {
	resultData, err := roles.GetRole()
	if err != nil {
		response.Context(ctx).Error(10002, err.Error())
		return
	}
	response.Context(ctx).Success(resultData)
	return
}

func RoleCreate(ctx *gin.Context) {
	var request roleRequestData
	if err := ctx.ShouldBind(&request); err != nil {
		response.Context(ctx).Error(10000, err.Error())
		return
	}
	var roleModel roles.Entity
	roleModel.RoleName = request.RoleName
	roleModel.RoleDesc = request.RoleDesc
	roleModel.Status = request.Status

	_, err := db.OrmClient().InsertOne(&roleModel)
	if err != nil {
		response.Context(ctx).Error(10001, err.Error())
		return
	}
	response.Context(ctx).Success(request)
	return
}

func RoleDetail(ctx *gin.Context) {
	queryId := ctx.Query("id")
	id, _ := strconv.Atoi(queryId)
	result, err := roles.GetRoleById(id)
	if err != nil {
		response.Context(ctx).Error(10000, err.Error())
		return
	}
	if result.Id == 0 {
		response.Context(ctx).Error(10001, "数据不存在")
		return
	}
	response.Context(ctx).Success(result)
	return
}

func RoleUpdate(ctx *gin.Context) {
	queryId := ctx.Query("id")
	id, _ := strconv.Atoi(queryId)
	var request roleRequestData
	if err := ctx.ShouldBind(&request); err != nil {
		response.Context(ctx).Error(10000, err.Error())
		return
	}
	var roleModel roles.Entity
	roleModel.RoleName = request.RoleName
	roleModel.RoleDesc = request.RoleDesc
	roleModel.Status = request.Status

	_, err := db.OrmClient().ID(id).Update(&roleModel)
	if err != nil {
		response.Context(ctx).Error(10001, err.Error())
		return
	}
	response.Context(ctx).Success(request)
	return
}

func RoleDelete(ctx *gin.Context) {
	queryId := ctx.Query("id")
	id, _ := strconv.Atoi(queryId)
	var role roles.Entity
	roleInfo, err := db.OrmClient().ID(id).Get(&role)
	if err != nil {
		response.Context(ctx).Error(10000, "获取数据错误"+err.Error())
		return
	}
	if !roleInfo {
		response.Context(ctx).Error(10001, "数据不存在")
		return
	}
	result, delErr := db.OrmClient().ID(id).Delete(&role)
	if delErr != nil {
		response.Context(ctx).Error(10002, "删除数据失败")
		return
	}
	response.Context(ctx).Success(result)
	return
}
