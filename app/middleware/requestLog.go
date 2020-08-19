package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Request.URL)
	}
}