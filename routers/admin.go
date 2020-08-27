package routers

import (
	"gin-api/app/controller/admin"
	"gin-api/app/middleware"
	"github.com/gin-gonic/gin"
)

// 后台路由组
func InitAdminRouter(router *gin.Engine) {
	router.POST("/admin/login", admin.Login)
	adminGroup := router.Group("/admin")
	{
		// 权限验证中间件
		adminGroup.Use(middleware.AdminAuth())
		adminGroup.POST("/detail", admin.Login)
		roleGroup := adminGroup.Group("/role")
		{
			roleGroup.POST("/list", admin.RoleList)
			roleGroup.POST("/detail", admin.RoleDetail)
			roleGroup.POST("/create", admin.RoleCreate)
			roleGroup.POST("/update", admin.RoleUpdate)
			roleGroup.POST("/delete", admin.RoleDelete)
		}
	}
}
