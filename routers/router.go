package routers

import (
	"gin-api/app/controller"
	"gin-api/app/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./public/assets")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	router.Use(middleware.RequestLog())

	router.GET("/", controller.Index)
}
