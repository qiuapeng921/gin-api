package response

import (
	"fmt"
	"gin-api/app/consts"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Wrapper struct {
	*gin.Context
}

func Context(c *gin.Context) *Wrapper {
	return &Wrapper{c}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response setting gin.JSON
func (wrapper *Wrapper) Response(httpCode, errCode int, data interface{}) {
	wrapper.JSON(httpCode, Response{
		Code:    errCode,
		Message: consts.GetMsg(errCode),
		Data:    data,
	})
	return
}

func (wrapper *Wrapper) View(name string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	wrapper.HTML(200, fmt.Sprintf("%s.html", name), responseData)
	return
}

func (wrapper *Wrapper) Success(data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	wrapper.Response(http.StatusOK, http.StatusOK, responseData)
	return
}

func (wrapper *Wrapper) Error(errCode int) {
	wrapper.Response(http.StatusOK, errCode, "")
	return
}