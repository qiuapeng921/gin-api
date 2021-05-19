package routers

import (
	"gin-api/app/http/controller"
	"gin-api/app/websocket"
	v1 "gin-api/routers/v1"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	router.GET("/", controller.Index)

	// 初始化v1路由组
	v1.Router(router.Group("/v1"))

	router.GET("/ws", func(context *gin.Context) {
		websocket.NewWebsocket(context)
	})
}
