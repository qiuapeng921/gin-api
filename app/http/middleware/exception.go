package middleware

import (
	"bytes"
	"fmt"
	"gin-api/app/utility/app"
	"gin-api/app/utility/response"
	"gin-api/app/utility/system"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type errorInfo struct {
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Method   string `json:"method"`
}

func HandleException() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				notice := os.Getenv("QY_NOTICE_KEY")
				if notice != "" {
					// 定义 文件名、行号、方法名
					fileName, line, functionName := "", 0, ""
					var pc uintptr
					var ok bool
					pc, fileName, line, ok = runtime.Caller(5)
					if ok {
						functionName = runtime.FuncForPC(pc).Name()
						functionName = filepath.Ext(functionName)
						functionName = strings.TrimPrefix(functionName, ".")
					}

					msg := errorInfo{fileName, line, functionName}

					errorJsonInfo := system.StructToJson(msg)

					content := fmt.Sprintf("### %s项目出错了 \n ###### 请求时间:%s \n ###### 请求地址:%s \n ###### ip地址:%s \n ###### 文件名及方法:%s \n ###### 错误信息:%v",
						"api-gin",
						system.GetCurrentDate(),
						c.Request.Method+"======"+c.Request.Host+c.Request.RequestURI,
						c.ClientIP(),
						errorJsonInfo,
						err,
					)

					message := gin.H{"msgtype": "markdown", "markdown": gin.H{"content": content, "mentioned_list": []string{"@all"}}}

					pushMessage := system.MapToJson(message)

					endPoint := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + notice
					_, err := app.Http().Post(endPoint, "application/json", bytes.NewBuffer([]byte(pushMessage)))
					if err != nil {
						fmt.Println(err.Error())
						return
					}
				}

				response.Context(c).String(http.StatusInternalServerError, "系统异常，请联系管理员！")
				c.Abort()
			}
		}()
		c.Next()
	}
}
