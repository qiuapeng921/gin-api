package roles

import "gin-api/helpers/pool/grom"

func GetRole() (entity []Entity, err error) {
	err = grom.GetOrm().Find(&entity)
	return
}

func GetRoleById(id int) (entity Entity, err error) {
	_, err = grom.GetOrm().Where("id=?", id).Get(&entity)
	return
}
