package controller

import (
	"gin-api/helpers/response"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.Context(c).View("index", gin.H{"name": "GinApi"})
	return
}

