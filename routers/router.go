package routers

import (
	"gin-api/app/http/middleware"
	"gin-api/app/socket"
	"gin-api/app/utility/response"
	"gin-api/app/utility/templates"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(router *gin.Engine) {
	// 加载模板引擎
	templates.InitTemplate(router)

	router.Use(
		middleware.Cors(),
		middleware.RequestId(),
	)

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/auth/login")
	})

	router.GET("/ws", func(ctx *gin.Context) {
		socket.Handler(ctx)
	})

	// 404错误
	router.NoRoute(func(ctx *gin.Context) {
		response.Context(ctx).View("error", gin.H{"message": "路由异常"})
		return
	})

	router.NoMethod(func(ctx *gin.Context) {
		response.Context(ctx).View("error", gin.H{"message": "请求方式错误"})
		return
	})

	// 加载后台路由组
	InitAdminRouter(router)
}
