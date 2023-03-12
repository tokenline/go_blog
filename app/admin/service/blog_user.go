package service

import (
	"go-blog/app/admin/dto"
	"go-blog/common/global"
	"go-blog/model"
)

type BlogUserService struct{}

// 添加用户
func (*BlogUserService) Add(addDto dto.BlogUserAddDto) error {

	var user = &model.BlogUser{
		Username: addDto.Username,
		Nickname: addDto.Nickname,
		Mail:     addDto.Mail,
		Phone:    addDto.Phone,
		Password: addDto.Password,
		Gender:   addDto.Gender,
		Status:   addDto.Status,
		Remark:   addDto.Remark,
	}
	//添加
	err := global.DB.Create(user).Error
	// if err!=nil{
	// 	panic("")
	// }
	return err
}

// 更新用户信息
func (*BlogUserService) Update(updateDto dto.BlogUserUpdateDto) error {

	var user = &model.BlogUser{}

	err := global.DB.Model(&user).Where("id = ?", updateDto.ID).Updates(map[string]interface{}{
		"nickname": updateDto.Nickname,
		"mail":     updateDto.Mail,
		"phone":    updateDto.Phone,
		"gender":   updateDto.Gender,
		"status":   updateDto.Status,
		"remark":   updateDto.Remark,
	}).Error

	return err
}
