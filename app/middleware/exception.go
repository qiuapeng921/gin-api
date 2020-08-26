package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-api/helpers/response"
	"gin-api/helpers/system"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
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
				// 定义 文件名、行号、方法名
				fileName, line, functionName := "?", 0, "?"
				var pc uintptr
				var ok bool
				pc, fileName, line, ok = runtime.Caller(5)
				if ok {
					functionName = runtime.FuncForPC(pc).Name()
					functionName = filepath.Ext(functionName)
					functionName = strings.TrimPrefix(functionName, ".")
				}

				msg := errorInfo{fileName, line, functionName}
				jsons, errs := json.Marshal(msg)
				if errs != nil {
					fmt.Println("json marshal error:", errs)
				}
				errorJsonInfo := string(jsons)

				content := fmt.Sprintf("### %s项目出错了 \n ###### 请求时间:%s \n ###### 请求地址:%s \n ###### ip地址:%s \n ###### 文件名及方法:%s \n ###### 错误信息:%v",
					"api-gin",
					system.GetCurrentDate(),
					c.Request.Method+"======"+c.Request.Host+c.Request.RequestURI,
					c.ClientIP(),
					errorJsonInfo,
					err,
				)

				message := gin.H{"msgtype": "markdown", "markdown": gin.H{"content": content, "mentioned_list": []string{"@all"}}}

				client := &http.Client{}
				endPoint := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=e0544aaf-ffa2-4ec0-b3c4-316fb2fcd8b6"
				req, err := client.Post(endPoint, "application/json", bytes.NewBuffer([]byte(system.MapToJson(message))))
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				result, _ := ioutil.ReadAll(req.Body)
				fmt.Println(string(result))

				response.Context(c).String(http.StatusInternalServerError, "系统异常，请联系管理员！")
				c.Abort()
			}
		}()
		c.Next()
	}
}
