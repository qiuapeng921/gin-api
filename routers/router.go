package routers

import (
	"gin-api/app/http/controller"
	"gin-api/app/websocket"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

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

	router.GET("/ws", func(context *gin.Context) {
		websocket.NewWebsocket(context)
	})
}
