package controller

import (
	"fmt"
	"gin-api/app/http/request"
	"gin-api/app/service"
	"gin-api/app/utility/app"
	"gin-api/app/utility/captcha"
	"gin-api/app/utility/response"
	"github.com/gin-gonic/gin"
)


// @生成验证码
// @Author 邱阿朋
// @Date 16:43 2020/11/04
func Captcha(ctx *gin.Context) {
	captcha.GenerateCaptcha(ctx)
	return
}

func Login(ctx *gin.Context) {
	response.Context(ctx).View("login")
	return
}


// @用户登录
// @Author 邱阿朋
// @Date 16:43 2020/11/04
func HandleLogin(ctx *gin.Context) {
	var params request.AuthRequest
	if err := ctx.ShouldBind(&params); err != nil {
		response.Context(ctx).Error(10000, err.Error())
		return
	}
	if !captcha.VerifyCaptcha(ctx, params.Code) {
		response.Context(ctx).Error(10001, "验证码错误")
		return
	}
	result, code, err := service.HandelAdminAuth(params.Username, params.Password)
	if err != nil {
		response.Context(ctx).Error(code, err.Error())
	} else {
		response.Context(ctx).Success(result)
	}
	return
}

// @Description
// @Author 邱阿朋
// @Date 16:43 2020/11/04
func Logout(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	app.Redis().Del(fmt.Sprintf("token:%d", id))
	response.Context(ctx).Success()
	return
}
