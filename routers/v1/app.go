package v1

import (
	"gin-api/app/http/controller"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	esGroup := router.Group("/es")
	{
		esGroup.GET("/list", controller.SearchList)
		esGroup.GET("/query", controller.SearchQuery)
		esGroup.GET("/create", controller.SearchCreate)
		esGroup.GET("/info", controller.SearchInfo)
		esGroup.GET("/update", controller.SearchUpdate)
		esGroup.GET("/delete", controller.SearchDelete)
	}
}
