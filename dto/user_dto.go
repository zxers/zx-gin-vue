package dto

import "github.com/zxers/zx-gin-vue/model"

type UserDto struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name: 	user.Name,
		Phone: 	user.Phone,
	}
}