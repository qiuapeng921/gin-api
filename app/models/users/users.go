package users

import "gin-api/helpers/db"

// 通过用户名获取用户
func GetUserByName(name string) (user Entity, err error) {
	_, err = db.Xorm().Where("username=?", name).Get(&user)
	return user, err
}
