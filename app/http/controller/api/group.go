package api

import (
	"gin-api/app/utility/response"
	"github.com/gin-gonic/gin"
)

// 创建分组
func CreateGroup(ctx *gin.Context) {
	response.Context(ctx).Success(gin.H{"data": "创建分组"})
	return
}

// 加入分组
func JoinGroup(ctx *gin.Context)  {
	response.Context(ctx).Success(gin.H{"data": "加入分组"})
	return
}

// 移除分组
func RemoveGroup(ctx *gin.Context)  {
	response.Context(ctx).Success(gin.H{"data": "移除分组"})
	return
}

// 删除分组
func DeleteGroup(ctx *gin.Context) {
	response.Context(ctx).Success(gin.H{"data": "删除分组"})
	return
}
