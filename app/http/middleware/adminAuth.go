package middleware

import (
	"gin-api/app/utility/jwt"
	"gin-api/app/utility/response"
	"github.com/gin-gonic/gin"
	"time"
)

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.URL.Path
		if uri != "/admin/login" {
			token := ctx.Request.Header.Get("token")
			if token == "" {
				response.Context(ctx).Error(20000, "token不能为空")
				return
			}
			result, err := jwt.ParseToken(token)
			if err != nil {
				response.Context(ctx).Error(20001, err.Error())
				return
			}
			if result.ExpiresAt < time.Now().Unix() {
				response.Context(ctx).Error(20002, "token过期")
				return
			}
			if result.Category != "admin" {
				response.Context(ctx).Error(20003, "token类型错误")
				return
			}
			ctx.Set("id", int(result.Id))
		}
		ctx.Next()
		return
	}
}
