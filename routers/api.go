package routers

import (
	"gin-api/app/controller/api"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(router *gin.Engine) {

	apiGroup := router.Group("/api")
	{
		userGroup := apiGroup.Group("/user")
		{
			userGroup.GET("/get_all", api.UserList)
			userGroup.GET("/insert", api.UserInsert)
			userGroup.GET("/update", api.UserUpdate)
			userGroup.GET("/delete", api.UserDelete)
			userGroup.GET("/force_delete", api.UserForceDelete)
		}
	}
}
