package controller

import (
	"gin-api/app/models/admins"
	"gin-api/app/utility/app"
	"gin-api/app/utility/response"
	"github.com/gin-gonic/gin"
)

// @获取列表页面
// @Author 邱阿朋
// @Date 12:29 2020/11/05
func GetPage(ctx *gin.Context) {
	response.Context(ctx).View("admin/list")
	return
}

func GetAdminList(ctx *gin.Context) {
	var admin admins.Entity
	_, _ = app.DB().Get(&admin)
	count, _ := app.DB().Count(&admin)
	response.Context(ctx).Page(int(count),admin)
	return
}