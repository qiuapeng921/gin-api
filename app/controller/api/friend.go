package api

import (
	"gin-api/helpers/response"
	"github.com/gin-gonic/gin"
)

// 添加好友关系
func CreateFriend(ctx *gin.Context)  {
	response.Context(ctx).Success(gin.H{"data": "创建好友关系"})
	return
}

// 删除好友关系
func DeleteFriend(ctx *gin.Context)  {
	response.Context(ctx).Success(gin.H{"data": "删除好友关系"})
	return
}
