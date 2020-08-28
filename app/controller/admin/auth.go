package admin

import (
	"gin-api/app/models/admins"
	"gin-api/helpers/db"
	"gin-api/helpers/jwt"
	"gin-api/helpers/response"
	"gin-api/helpers/system"
	"github.com/gin-gonic/gin"
	"time"
)

type authRequestData struct {
	UserName string `json:"username" from:"username" binding:"required"`
	Password string `json:"password" from:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var request authRequestData
	if err := ctx.ShouldBind(&request); err != nil {
		response.Context(ctx).Error(10000, err.Error())
		return
	}

	result, err := admins.GetAdminByUserName(request.UserName)
	if err != nil {
		response.Context(ctx).Error(10001, err.Error())
		return
	}
	if result.Id == 0 {
		response.Context(ctx).Error(10002, "用户未找到")
		return
	}
	if system.EncodeMD5(request.Password) != result.Password {
		response.Context(ctx).Error(10003, "密码错误")
		return
	}
	token, expiresAt, genTokenErr := jwt.GenerateToken(uint(result.Id), result.Username, "admin")
	tokenExpiresAt := time.Now().Unix()

	_, cacheErr := db.Redis().Set("admin_token:"+result.Username, token, time.Duration(expiresAt-tokenExpiresAt)*time.Second).Result()
	if cacheErr != nil {
		response.Context(ctx).Error(10003, "cache err"+cacheErr.Error())
		return
	}
	if genTokenErr != nil {
		response.Context(ctx).Error(10004, genTokenErr.Error())
		return
	}
	response.Context(ctx).Success(gin.H{
		"user":        result,
		"token":       token,
		"expireAt":    expiresAt,
		"permissions": gin.H{"id": "queryForm", "operation": []string{"add", "edit", "delete"}},
		"roles":       gin.H{"id": "admin", "operation": []string{"add", "edit", "delete"}},
	})
	return
}