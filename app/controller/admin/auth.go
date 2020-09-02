package admin

import (
	"gin-api/app/service"
	"gin-api/helpers/response"
	"github.com/gin-gonic/gin"
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
	result, code, err := service.HandelAdminAuth(request.UserName, request.Password)
	if err != nil {
		response.Context(ctx).Error(code, err.Error())
	} else {
		response.Context(ctx).Success(result)
	}
	return
}
