package roles

import (
	"gin-api/helpers/db"
)

func GetRole() (entity []Entity, err error) {
	err = db.Xorm().Find(&entity)
	return
}

func GetRoleById(id int) (entity Entity, err error) {
	_, err = db.Xorm().Where("id=?", id).Get(&entity)
	return
}
