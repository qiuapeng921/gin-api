package routers

import (
	"gin-api/app/http/controller/admin"
	"gin-api/app/http/middleware"
	"github.com/gin-gonic/gin"
)

// 后台路由组
func InitAdminRouter(router *gin.Engine) {
	adminGroup := router.Group("/admin")
	{
		adminGroup.POST("/login", admin.Login)

		// 权限验证中间件 中间件上面的不做token验证
		adminGroup.Use(middleware.AdminAuth())
	}
}
