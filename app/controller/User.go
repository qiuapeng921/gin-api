package controller

import (
	"fmt"
	"gin-api/app/models/users"
	"gin-api/helpers/pool/grom"
	"gin-api/helpers/response"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context)  {
	var userModel []users.Entity
	err := grom.GetOrm().Unscoped().OrderBy("id desc").Find(&userModel)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	response.Context(c).Success(userModel)
	return
}

func UserInsert(c *gin.Context)  {
	for i := 0; i < 10; i++ {
		var userModel users.Entity
		userModel.Username = "admin" + fmt.Sprintf("%d", i)
		userModel.Password = "123456"
		userModel.Phone = "1524927977" + fmt.Sprintf("%d", i)
		result, err := grom.GetOrm().Insert(&userModel)
		if err != nil {
			fmt.Println("err", err.Error())
		}
		fmt.Println(result)
	}
	return
}

func UserUpdate(c *gin.Context)  {
	avatar := c.Query("avatar")
	var userModel users.Entity
	userModel.Avatar = avatar
	result, err := grom.GetOrm().Where("status=?", 0).Update(&userModel)
	if err != nil {
		response.Context(c).JSON(200, gin.H{"err": err.Error()})
		return
	}
	response.Context(c).Success(result)
	return
}

func UserDelete(c *gin.Context)  {
	var userModel users.Entity
	result, err := grom.GetOrm().Where("status=?", 0).Delete(&userModel)
	if err != nil {
		response.Context(c).JSON(200, gin.H{"err": err.Error()})
		return
	}
	response.Context(c).Success(result)
	return
}

func UserForceDelete(c *gin.Context)  {
	var userModel users.Entity
	result, err := grom.GetOrm().Where("status=?", 0).Unscoped().Delete(&userModel)
	if err != nil {
		response.Context(c).JSON(200, gin.H{"err": err.Error()})
		return
	}
	response.Context(c).Success(result)
	return
}