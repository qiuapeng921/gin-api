package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RequestLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Printf("域名：%s ,请求地址：%s \n",
			ctx.Request.Host,
			ctx.Request.URL.Path,
		)
		ctx.Next()
	}
}
