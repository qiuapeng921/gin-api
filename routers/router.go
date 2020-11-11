package routers

import (
	"gin-api/app/http/controller"
	"gin-api/app/http/middleware"
	"gin-api/app/socket"
	"gin-api/app/utility/templates"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	templates.InitTemplate(router)

	router.Use(
		middleware.RequestLog(),
		middleware.Cors(),
		middleware.RequestId(),
	)

	router.GET("/", controller.Index)

	esGroup := router.Group("/es")
	{
		esGroup.GET("/list", controller.SearchList)
		esGroup.GET("/query", controller.SearchQuery)
		esGroup.GET("/create", controller.SearchCreate)
		esGroup.GET("/info", controller.SearchInfo)
		esGroup.GET("/update", controller.SearchUpdate)
		esGroup.GET("/delete", controller.SearchDelete)
	}

	router.GET("/ws", socket.Handler)

	// 加载后台路由组
	InitAdminRouter(router)
	// 加载前台路由组
	InitApiRouter(router)
}
