package controller

import (
	"gin-api/helpers/db"
	"gin-api/helpers/response"
	"gin-api/helpers/system"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func Index(c *gin.Context) {
	response.Context(c).View("index", gin.H{"name": "GinApi"})
	return
}

func Elastic(ctx *gin.Context) {
	index := ctx.Query("index")
	query := ctx.Query("query")
	data := system.MapToJson(gin.H{"username": GetRandomString(10), "password": "123456"})
	putData := db.EsClient.PutData(index, data)
	queryResult := db.EsClient.Query(index)
	list := db.EsClient.List(index, map[string]string{"query": query})
	response.Context(ctx).Success(gin.H{
		"putData": putData,
		"query":   queryResult,
		"list":    list,
	})
	return
}
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
