package response

import (
	"fmt"
	"gin-admin/app/consts"
	"github.com/gin-gonic/gin"
	"net/http"
)

type wrapper struct {
	*gin.Context
}

func Context(c *gin.Context) *wrapper {
	return &wrapper{c}
}

type response struct {
	Code    int         `json:"code"`
	Count   int         `json:"count"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (wrapper *wrapper) View(name string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	wrapper.HTML(http.StatusOK, fmt.Sprintf("%s.html", name), responseData)
	wrapper.Abort()
	return
}

func (wrapper *wrapper) Success(data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	wrapper.JSON(http.StatusOK, response{
		Code:    0,
		Status:  http.StatusOK,
		Message: "Success",
		Data:    responseData,
	})
	wrapper.Abort()
	return
}

func (wrapper *wrapper) Page(count int, data ...interface{}) {
	wrapper.JSON(http.StatusOK, response{
		Code:    0,
		Count:   count,
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	})
	wrapper.Abort()
	return
}

func (wrapper *wrapper) Error(errCode int, message ...string) {
	responseMessage := consts.GetMsg(errCode)
	if len(message) > 0 {
		responseMessage = message[0]
	}
	wrapper.JSON(http.StatusOK, response{
		Code:    0,
		Status:  errCode,
		Message: responseMessage,
	})
	wrapper.Abort()
	return
}
