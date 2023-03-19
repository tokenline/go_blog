package service

import (
	"errors"
	"go-blog/app/admin/dto"
	"go-blog/common/global"
	"go-blog/constants"
	"go-blog/model"
)

type BlogAuthService struct{}

func (*BlogAuthService) Login(login dto.LoginDto) error {
	//获取user信息
	var user model.BlogUser
	err := global.DB.Where("username = ?", login.Username).First(&user).Error

	//验证用户是否存在
	if err != nil {
		return err
	}

	//用户禁用，返回错误
	if user.Status == constants.UserStatusDisable {
		return errors.New("账号已被禁用")
	}

	//验证用户密码是否正确
	if user.Password != login.Password {
		return errors.New("密码错误")
	}
	return err

}
