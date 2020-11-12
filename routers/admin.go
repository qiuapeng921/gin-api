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
		authGroup.POST("login", controller.Login)

		authGroup.GET("captcha", controller.Captcha)

		// 权限验证中间件 中间件上面的不做token验证
		//authGroup.Use(middleware.AdminAuth())

		authGroup.GET("logout", controller.Logout)
	}

	groups := router.Group("admin")
	{
		// 权限验证中间件 中间件上面的不做token验证
		groups.Use(middleware.AdminAuth())
		groups.GET("index", controller.Index)
		groups.GET("dashboard", controller.Dashboard)

		groups.GET("page", controller.GetPage)
		groups.Any("list", controller.GetAdminList)
	}
}