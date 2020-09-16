package api

import (
	"gin-api/app/models/users"
	"gin-api/app/utility/response"
	"github.com/gin-gonic/gin"
)

func Detail(ctx *gin.Context) {
	id := ctx.GetInt("id")
	result, resErr := users.GetUserById(id)
	if resErr != nil {
		response.Context(ctx).Error(10001, resErr.Error())
		return
	}
	response.Context(ctx).Success(result)
	return
}
