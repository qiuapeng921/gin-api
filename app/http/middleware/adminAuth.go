package middleware

import (
	"gin-api/app/utility/auth"
	"gin-api/app/utility/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.URL.Path
		if uri != "/auth/login" {
			token, err := ctx.Request.Cookie("access_token")
			if err == http.ErrNoCookie || token == nil || token.Value == "" {
				ctx.Redirect(http.StatusMovedPermanently, "/auth/login")
				ctx.Abort()
				return
			}
			result, err := auth.ParseToken(token.Value)
			if err != nil {
				response.Context(ctx).View("error", gin.H{"message": err.Error()})
				ctx.Abort()
				return
			}
			if result.ExpiresAt < time.Now().Unix() {
				response.Context(ctx).View("error", gin.H{"message": "token过期"})
				ctx.Abort()
				return
			}
			if result.Category != "admin" {
				response.Context(ctx).View("error", gin.H{"message": "token类型错误"})
				ctx.Abort()
				return
			}
			ctx.Set("id", int(result.Id))
		}
		ctx.Next()
	}
}
