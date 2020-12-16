package websocket

import (
	"gin-api/app/websocket/server"
	"github.com/gin-gonic/gin"
)

func NewWebsocket(ctx *gin.Context) {
	socketServer := server.NewWebSocketServer()
	socketServer.Upgrade(ctx.Writer, ctx.Request)
}