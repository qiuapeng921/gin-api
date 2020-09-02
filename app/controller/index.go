package controller

import (
	"gin-api/helpers/db"
	"gin-api/helpers/response"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.Context(c).View("index", gin.H{"name": "GinApi"})
	return
}

func Elastic(ctx *gin.Context) {
	index := ctx.Query("index")
	query := db.EsClient.Query(index)
	list := db.EsClient.List(index, map[string]string{"query": "admin1"})
	response.Context(ctx).Success(gin.H{
		"query": query,
		"list":  list,
	})
	return
}
