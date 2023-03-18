package dto

import "go-blog/constants"

type BlogUserAddDto struct {
	Username string               //用户名
	Nickname string               //昵称
	Mail     string               //邮箱地址
	Phone    string               //电话号码
	Password string               //密码
	Gender   constants.UserGender //性别
	Status   constants.UserStatus //状态
	Remark   string               //备注
}

type BlogUserUpdateDto struct {
	ID       uint //用户ID  想一想
	Nickname string
	Mail     string
	Phone    string
	Gender   constants.UserGender
	Status   constants.UserStatus
	Remark   string
}

// 删除应该只需要一个ID

type BlogUserVo struct {
	ID       uint                 //用户ID
	Username string               //用户名
	Nickname string               //昵称
	Mail     string               //邮箱地址
	Phone    string               //电话号码
	Gender   constants.UserGender //性别
	Status   constants.UserStatus //状态
	Remark   string               //备注
}

// TODO:换公共Dto
type BlogUserDeleteDto struct {
	ID uint
}

// 这个应该就是一个ID吧？
type BlogUserQueryDto struct {
	ID uint //用户ID
}
