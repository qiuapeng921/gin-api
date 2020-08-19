package admins

import "gin-api/helpers/pool/grom"

func GetAdminById(id int) (res bool, err error, entity Entity) {
	res, err = grom.GetOrm().Where("id=?", id).Get(&entity)
	return
}

func GetAdminByUserName(username string) (res bool, err error) {
	res, err = grom.GetOrm().Where("username=?", username).Get(&Entity{})
	return
}
