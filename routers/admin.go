package routers

import (
	"gin-api/app/http/controller"
	"gin-api/app/http/middleware"
	"github.com/gin-gonic/gin"
)

// 后台路由组
func InitAdminRouter(router *gin.Engine) {

	authGroup := router.Group("auth")
	{
		authGroup.GET("login", controller.Login)
		authGroup.POST("handle_login", controller.HandleLogin)
		authGroup.GET("captcha", controller.Captcha)
		authGroup.GET("logout", controller.Logout).Use(middleware.AdminAuth())
	}

	groups := router.Group("admin").Use(middleware.AdminAuth())
	{
		groups.GET("index", controller.Index)
		groups.GET("dashboard", controller.Dashboard)

		groups.GET("page", controller.GetPage)
		groups.GET("list", controller.GetAdminList)
	}
}