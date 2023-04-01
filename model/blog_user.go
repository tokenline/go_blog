package model

import (
	"go-blog/constants"

	"gorm.io/gorm"
)

type BlogUser struct {
	gorm.Model
	Username string               //用户名
	Nickname string               //昵称
	Mail     string               //邮箱地址
	Phone    string               //电话号码
	Password string               //密码
	Salt     string               //密码盐
	Gender   constants.UserGender //性别
	Status   constants.UserStatus //状态
	Remark   string               //备注
}
