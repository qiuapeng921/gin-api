package admins

import (
	"gin-api/app/utility/app"
)

func GetAdminById(id int) (entity Entity, err error) {
	_, err = app.DB().Where("id = ?", id).Get(&entity)
	return
}

func GetAdminByUserName(username string) (entity Entity, err error) {
	_, err = app.DB().Where("username = ?", username).Get(&entity)
	return
}
