package users

import (
	"gin-api/app/utility/app"
)

func GetUserById(id int) (user Entity, err error) {
	_, err = app.DB().ID(id).Get(&user)
	return user, err
}

// 通过用户名获取用户
func GetUserByName(name string) (user Entity, err error) {
	_, err = app.DB().Where("username=?", name).Get(&user)
	return user, err
}
