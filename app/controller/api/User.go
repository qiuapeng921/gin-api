package api

import (
	"fmt"
	"gin-api/app/models/users"
	"gin-api/helpers/db"
	"gin-api/helpers/response"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	var userModel []users.Entity
	err := db.Xorm().Unscoped().OrderBy("id desc").Find(&userModel)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	response.Context(c).Success(userModel)
	return
}

func UserInsert(c *gin.Context) {
	var responseData []int64
	var errCount []int64
	for i := 0; i < 10; i++ {
		var userModel users.Entity
		userModel.Username = "admin" + fmt.Sprintf("%d", i)
		userModel.Password = "123456"
		userModel.Phone = "1524927977" + fmt.Sprintf("%d", i)
		result, err := db.Xorm().Insert(&userModel)
		if err != nil {
			errCount = append(errCount, result)
		} else {
			responseData = append(responseData, result)
		}
	}
	response.Context(c).Success(gin.H{"success": len(responseData), "error": len(errCount)})
	return
}

func UserUpdate(c *gin.Context) {
	avatar := c.Query("avatar")
	var userModel users.Entity
	userModel.Avatar = avatar
	result, err := db.Xorm().Where("status=?", 0).Update(&userModel)
	if err != nil {
		response.Context(c).JSON(200, gin.H{"err": err.Error()})
		return
	}
	response.Context(c).Success(result)
	return
}

func UserDelete(c *gin.Context) {
	var userModel users.Entity
	result, err := db.Xorm().Where("status=?", 0).Delete(&userModel)
	if err != nil {
		response.Context(c).JSON(200, gin.H{"err": err.Error()})
		return
	}
	response.Context(c).Success(result)
	return
}

func UserForceDelete(c *gin.Context) {
	var userModel users.Entity
	result, err := db.Xorm().Where("status=?", 0).Unscoped().Delete(&userModel)
	if err != nil {
		response.Context(c).JSON(200, gin.H{"err": err.Error()})
		return
	}
	response.Context(c).Success(result)
	return
}
