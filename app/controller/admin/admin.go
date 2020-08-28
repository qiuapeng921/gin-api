package admin

import (
	"gin-api/app/models/admins"
	"gin-api/helpers/response"
	"github.com/gin-gonic/gin"
)


func Detail(ctx *gin.Context) {
	id := ctx.GetInt("id")
	result, resErr := admins.GetAdminById(id)
	if resErr != nil {
		response.Context(ctx).Error(10001, resErr.Error())
		return
	}
	response.Context(ctx).Success(result)
	return
}
