package model

import (
	"go-blog/constant"

	"gorm.io/gorm"
)

type BlogUser struct {
	gorm.Model
	Username string              //用户名
	Nickname string              //昵称
	Mail     string              //邮箱地址
	Phone    string              //电话号码
	Password string              //密码
	Gender   constant.UserGender //性别
	Status   int                 //状态
	Remark   string              //备注
}
