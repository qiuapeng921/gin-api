package controller

import (
	"gin-api/helpers/db"
	"gin-api/helpers/response"
	"gin-api/helpers/system"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

// 索引数据列表
// query admin
// page 1
// size 10
// sort _source|_index|_id
// sort_type desc|aec
func SearchList(ctx *gin.Context) {
	var result []*elastic.SearchHit
	if db.EsClient.IsExistsIndex("search") {
		result = db.EsClient.List("search", map[string]string{"size": "10"})
	}
	response.Context(ctx).Success(result)
	return
}

// 索引数据搜索
// query admin
func SearchQuery(ctx *gin.Context) {
	query := ctx.Query("query")
	var result []*elastic.SearchHit
	if db.EsClient.IsExistsIndex("search") {
		result = db.EsClient.Query("search", query)
	}
	response.Context(ctx).Success(result)
	return
}

// 索引数据创建
func SearchCreate(ctx *gin.Context) {
	data := system.MapToJson(gin.H{"username": system.GetRandomString(5), "password": system.GetRandomString(10)})
	putData := db.EsClient.Insert("search", data)
	response.Context(ctx).Success(putData)
	return
}

// 索引数据详情
// id baidu.com
func SearchInfo(ctx *gin.Context) {
	id := ctx.Query("id")
	result := db.EsClient.Gets("search", id)
	response.Context(ctx).Success(result)
	return
}

// 索引数据详情
// id baidu.com
// username admin
// password 123456
func SearchUpdate(ctx *gin.Context) {
	id := ctx.Query("id")
	username := ctx.Query("username")
	password := ctx.Query("password")
	result := db.EsClient.UpdateData("search", id, gin.H{"username": username, "password": system.EncodeMD5(password)})
	response.Context(ctx).Success(result)
	return
}

// 索引数据删除
// id baidu.com
func SearchDelete(ctx *gin.Context) {
	id := ctx.Query("id")
	result := db.EsClient.DeleteData("search", id)
	response.Context(ctx).Success(result)
	return
}
