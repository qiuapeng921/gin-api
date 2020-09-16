package roles

import (
	"gin-api/app/utility/db"
)

func GetRole() (entity []Entity, err error) {
	err = db.OrmClient().Find(&entity)
	return
}

func GetRoleById(id int) (entity Entity, err error) {
	_, err = db.OrmClient().Where("id=?", id).Get(&entity)
	return
}
