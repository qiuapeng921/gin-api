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

	router.Use(middleware.RequestLog(),middleware.Cors())

	router.GET("/", controller.Index)
	router.GET("/ws", controller.WebSocketHandler)


	api := router.Group("/api")
	{
		api.POST("/login", controller.Login)
	}

	user := router.Group("/user")
	{
		user.GET("/get_all", controller.UserList)
		user.GET("/insert", controller.UserInsert)
		user.GET("/update", controller.UserUpdate)
		user.GET("/delete", controller.UserDelete)
		user.GET("/force_delete", controller.UserForceDelete)
	}
}
