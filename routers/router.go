package routers

import (
	"gin-api/app/controller"
	"gin-api/app/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	router.Use(middleware.RequestLog(), middleware.Cors())

	router.GET("/", controller.Index)
	router.GET("/ws", controller.WebSocketHandler)

	// 加载后台路由组
	InitAdminRouter(router)
	// 加载前台路由组
	InitApiRouter(router)
}
