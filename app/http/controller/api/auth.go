package api

import (
	"gin-api/app/models/users"
	"gin-api/app/utility/db"
	"gin-api/app/utility/jwt"
	"gin-api/app/utility/response"
	"gin-api/app/utility/system"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username" form:"username" binding:"required" comment:"用户名"`
	Password string `json:"password" form:"password" binding:"required" comment:"密码"`
}

func Register(c *gin.Context) {
	var request loginRequest
	if err := c.ShouldBind(&request); err != nil {
		response.Context(c).Error(10000, err.Error())
		return
	}
	// 判断用户是否存在
	result, err := users.GetUserByName(request.Username)
	if err != nil {
		response.Context(c).Error(10001, err.Error())
		return
	}
	if result.Id > 0 {
		response.Context(c).Error(10002, "用户已被注册")
		return
	}
	result.Username = request.Username
	result.Password = system.EncodeMD5(request.Username)
	insertId, insertErr := db.OrmClient().InsertOne(result)
	if insertErr != nil {
		response.Context(c).Error(10002, "用户注册失败"+insertErr.Error())
		return
	}
	response.Context(c).Success(insertId)
	return
}

func Login(c *gin.Context) {
	var request loginRequest
	if err := c.ShouldBind(&request); err != nil {
		response.Context(c).Error(10000, err.Error())
		return
	}
	// 判断用户是否存在
	user, err := users.GetUserByName(request.Username)
	if err != nil {
		response.Context(c).Error(10001, err.Error())
		return
	}
	if user.Id == 0 {
		response.Context(c).Error(10002, "用户未找到")
		return
	}

	if system.EncodeMD5(request.Password) != user.Password {
		response.Context(c).Error(10003, "账号密码不匹配")
		return
	}
	token, time, jwtErr := jwt.GenerateToken(uint(user.Id), user.Username, "api")
	if jwtErr != nil {
		response.Context(c).Error(10004, "token生成失败"+jwtErr.Error())
		return
	}
	userResponse := make(map[string]interface{})
	userResponse["id"] = user.Id
	userResponse["username"] = user.Username
	userResponse["avatar"] = user.Avatar
	response.Context(c).Success(gin.H{"user": userResponse, "token": token, "time": time})
	return
}