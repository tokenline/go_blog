package service

import (
	"errors"
	"go-blog/app/admin/dto"
	"go-blog/common/global"
	"go-blog/common/utils"
	"go-blog/constants"
	"go-blog/model"
	"time"
)

type BlogAuthService struct{}

func (*BlogAuthService) Login(login dto.LoginDto) (string, error) {
	//获取user信息
	var user model.BlogUser
	err := global.DB.Where("username = ?", login.Username).First(&user).Error

	//验证用户是否存在
	if err != nil {
		return "", errors.New("账号不存在")
	}

	//用户禁用，返回错误
	if user.Status == constants.UserStatusDisable {
		return "", errors.New("账号已被禁用")
	}
	pwd := utils.Encryption.Md5(login.Password + user.Salt)

	//验证用户密码是否正确
	if pwd != user.Password {
		return "", errors.New("密码错误")
	}

	var claims = utils.UserAuthClaims{
		UserId: user.ID,
	}
	//生成token
	token, err := utils.Jwt.GenerateToken(claims, time.Now().AddDate(0, 0, 1))

	return token, err

}
