package controller

import (
	"fmt"
	"gin-admin/app/http/request"
	"gin-admin/app/service"
	"gin-admin/app/utility/app"
	"gin-admin/app/utility/captcha"
	"gin-admin/app/utility/response"
	"github.com/gin-gonic/gin"
	"net/http"
)


// @生成验证码
// @Author 邱阿朋
// @Date 16:43 2020/11/04
func Captcha(ctx *gin.Context) {
	captcha.GenerateCaptcha(ctx)
	return
}

// @用户登录
// @Author 邱阿朋
// @Date 16:43 2020/11/04
func Login(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodGet {
		response.Context(ctx).View("login")
		return
	}

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
