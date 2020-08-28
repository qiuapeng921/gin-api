package routers

import (
	"gin-api/app/controller/api"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(router *gin.Engine) {

	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/register", api.UserRegister)
		apiGroup.POST("/login", api.UserLogin)
	}
}