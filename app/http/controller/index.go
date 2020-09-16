package controller

import (
	"gin-api/app/utility/response"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.Context(c).View("index", gin.H{"name": "GinApi"})
	return
}