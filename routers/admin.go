package routers

import (
	"gin-admin/app/http/controller"
	"gin-admin/app/http/middleware"
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
		authGroup.Use(middleware.AdminAuth())
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
		groups.POST("detail", controller.GetAdminDetail)
		groups.POST("create", controller.CreateAdmin)
		groups.POST("update", controller.UpdateAdmin)
		groups.POST("delete", controller.DeleteAdmin)

		roleGroup := groups.Group("role")
		{
			roleGroup.POST("list", controller.RoleList)
			roleGroup.POST("detail", controller.RoleDetail)
			roleGroup.POST("create", controller.RoleCreate)
			roleGroup.POST("update", controller.RoleUpdate)
			roleGroup.POST("delete", controller.RoleDelete)
		}
	}
}