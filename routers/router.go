package routers

import (
	"gin-api/app/controller"
	"gin-api/app/middleware"
	"gin-api/app/socket"
	"gin-api/helpers/templates"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	templates.InitTemplate(router)

	router.Use(
		middleware.RequestLog(),
		middleware.Cors(),
		middleware.HandleException(),
	)

	router.GET("/", controller.Index)
	router.GET("/ws", socket.Handler)

	// 加载后台路由组
	InitAdminRouter(router)
	// 加载前台路由组
	InitApiRouter(router)
}
