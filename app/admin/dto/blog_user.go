package dto

import "go-blog/constant"

type BlogUserAddDto struct {
	Username string              //用户名
	Nickname string              //昵称
	Mail     string              //邮箱地址
	Phone    string              //电话号码
	Password string              //密码
	Gender   constant.UserGender //性别
	Status   constant.UserStatus //状态
	Remark   string              //备注
}

type BlogUserUpdateDto struct {
	ID       int //用户ID  想一想
	Nickname string
	Mail     string
	Phone    string
	Gender   constant.UserGender
	Status   constant.UserStatus
	Remark   string
}
