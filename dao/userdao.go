package dao

import (

	"github.com/zxers/zx-gin-vue/db"
	"github.com/zxers/zx-gin-vue/model"
)

func GetUserByPhone(phone string) (*model.User, error) {
	user := new(model.User)
	err := db.DB.Debug().Where("phone=?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func IsPhoneExist(phone string) bool {
	var user model.User
	db.DB.Where("phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}