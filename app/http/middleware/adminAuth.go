package middleware

import (
	jwt "gin-admin/app/utility/auth"
	"gin-admin/app/utility/response"
	"github.com/gin-gonic/gin"
	"time"
)

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.URL.Path
		if uri != "/auth/login" {
			cookie, _ := ctx.Request.Cookie("access_token")
			token := cookie.Value
			if token = cookie.Value; token == "" {
				response.Context(ctx).View("error", gin.H{"message": "未登录..."})
				return
			}
			result, err := jwt.ParseToken(token)
			if err != nil {
				response.Context(ctx).View("error", gin.H{"message": err.Error()})
				return
			}
			if result.ExpiresAt < time.Now().Unix() {
				response.Context(ctx).View("error", gin.H{"message": "token过期"})
				return
			}
			if result.Category != "admin" {
				response.Context(ctx).View("error", gin.H{"message": "token类型错误"})
				return
			}
			ctx.Set("id", int(result.Id))
		}
		ctx.Next()
		return
	}
}
