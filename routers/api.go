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
		friendGroup := apiGroup.Group("/friend")
		{
			friendGroup.POST("/create", api.CreateFriend)
			friendGroup.POST("/delete", api.DeleteFriend)
		}
		group := apiGroup.Group("/group")
		{
			group.POST("/create", api.CreateGroup)
			group.POST("/join", api.JoinGroup)
			group.POST("/remove", api.RemoveGroup)
			group.POST("/delete", api.DeleteGroup)
		}
	}
}
