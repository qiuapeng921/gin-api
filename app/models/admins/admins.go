package admins

import "gin-api/helpers/pool/grom"

func GetAdminById(id int) (entity Entity, err error) {
	_, err = grom.GetOrm().Where("id = ?", id).Get(&entity)
	return
}

func GetAdminByUserName(username string) (entity Entity, err error) {
	_, err = grom.GetOrm().Where("username = ?", username).Get(&entity)
	return
}
