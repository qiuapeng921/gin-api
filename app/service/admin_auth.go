package service

import (
	"errors"
	"gin-api/app/models/admins"
	"gin-api/helpers/db"
	"gin-api/helpers/jwt"
	"gin-api/helpers/system"
	"github.com/gin-gonic/gin"
	"time"
)

func HandelAdminAuth(username, password string) (interface{}, int, error) {
	admin, err := admins.GetAdminByUserName(username)
	if err != nil {
		return nil, 10001, err
	}
	if admin.Id == 0 {
		return nil, 10002, errors.New("用户未找到")
	}
	if system.EncodeMD5(password) != admin.Password {
		return nil, 10003, errors.New("密码错误")
	}

	token, expiresAt, genTokenErr := jwt.GenerateToken(uint(admin.Id), admin.Username, "admin")
	tokenExpiresAt := time.Now().Unix()

	_, cacheErr := db.RedisClient().Set("admin_token:"+admin.Username, token, time.Duration(expiresAt-tokenExpiresAt)*time.Second).Result()
	if cacheErr != nil {
		return nil, 10003, cacheErr
	}
	if genTokenErr != nil {
		return nil, 10004, genTokenErr
	}
	responseData := gin.H{
		"admin":  gin.H{"id": admin.Id, "username": admin.Username, "phone": admin.Phone},
		"token":  token,
		"expire": expiresAt - system.GetCurrentUnix(),
	}
	return responseData, 200, nil
}