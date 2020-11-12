package response

import (
	"fmt"
	"gin-api/app/consts"
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
	Code   int `json:"code"`
	Status int `json:"status"`
}

type pageResponse struct {
	response
	Count int `json:"count"`
	successResponse
}

type successResponse struct {
	response
	Data interface{} `json:"data"`
}

type errorResponse struct {
	response
	Message string `json:"message"`
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
	var successResponse successResponse
	successResponse.Status = http.StatusOK
	successResponse.Data = responseData
	wrapper.JSON(http.StatusOK, successResponse)
	wrapper.Abort()
	return
}

func (wrapper *wrapper) Page(count int, data ...interface{}) {
	var pageResponse pageResponse
	pageResponse.Status = http.StatusOK
	pageResponse.Data = data
	wrapper.JSON(http.StatusOK, pageResponse)
	wrapper.Abort()
	return
}

func (wrapper *wrapper) Error(errCode int, message ...string) {
	responseMessage := consts.GetMsg(errCode)
	if len(message) > 0 {
		responseMessage = message[0]
	}
	var errorResponse errorResponse
	errorResponse.Status = errCode
	errorResponse.Message = responseMessage
	wrapper.JSON(http.StatusOK, errorResponse)
	wrapper.Abort()
	return
}
