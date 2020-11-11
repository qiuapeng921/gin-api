package controller

import (
	"gin-admin/app/models/admins"
	"gin-admin/app/utility/response"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	queryId, _ := c.Get("id")
	admin, _ := admins.GetAdminById(queryId.(int))
	response.Context(c).View("index", gin.H{"admin": admin})
	return
}
func Dashboard(c *gin.Context) {
	response.Context(c).View("dashboard")
	return
}
