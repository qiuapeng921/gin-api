package admins

import (
	"gin-api/helpers/db"
)

func GetAdminById(id int) (entity Entity, err error) {
	_, err = db.Xorm().Where("id = ?", id).Get(&entity)
	return
}

func GetAdminByUserName(username string) (entity Entity, err error) {
	_, err = db.Xorm().Where("username = ?", username).Get(&entity)
	return
}
