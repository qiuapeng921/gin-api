package routers

import (
	"gin-api/app/controller/api"
	"gin-api/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(router *gin.Engine) {

	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/register", api.Register)
		apiGroup.POST("/login", api.Login)

		// 中间件上面的不做token验证
		apiGroup.Use(middleware.ApiAuth())
		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/detail", api.Detail)
		}
	}
}
