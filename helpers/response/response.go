package response

import (
	"fmt"
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

func (wrapper *Wrapper) View(name string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	wrapper.HTML(http.StatusOK, fmt.Sprintf("%s.html", name), responseData)
	return
}

func (wrapper *Wrapper) Success(data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	wrapper.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    responseData,
	})
	return
}

func (wrapper *Wrapper) Error(errCode int, message string) {
	wrapper.JSON(http.StatusOK, Response{
		Code:    errCode,
		Message: message,
	})
	return
}
