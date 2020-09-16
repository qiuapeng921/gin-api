package admins

import (
	"gin-api/app/utility/db"
)

func GetAdminById(id int) (entity Entity, err error) {
	_, err = db.OrmClient().Where("id = ?", id).Get(&entity)
	return
}

func GetAdminByUserName(username string) (entity Entity, err error) {
	_, err = db.OrmClient().Where("username = ?", username).Get(&entity)
	return
}
