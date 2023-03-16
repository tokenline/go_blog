package service

import (
	"fmt"
	"go-blog/app/admin/dto"
	"go-blog/common/global"
	"go-blog/model"
)

type BlogUserService struct{}

// 添加用户
func (b *BlogUserService) Add(addDto dto.BlogUserAddDto) error {

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
func (b *BlogUserService) Update(updateDto dto.BlogUserUpdateDto) error {

	var user = &model.BlogUser{}

	global.DB.Model(&user).Where("ID = ?", updateDto.ID).Find(&user)

	// 判断updateDto中空的选项，并填入原有user中的信息

	// 感觉写一堆if 好蠢

	if updateDto.Nickname == "" {
		updateDto.Nickname = user.Nickname
	}

	if updateDto.Mail == "" {
		updateDto.Mail = user.Mail
	}

	if updateDto.Phone == "" {
		updateDto.Phone = user.Phone
	}

	// 性别和状态都遇到了一样的问题，这里如何表示nil?
	// if updateDto.Gender == nil {
	// 	updateDto.Gender = user.Gender
	// }

	// if updateDto.Status == nil {
	// 	updateDto.Status = user.Status
	// }

	if updateDto.Remark == "" {
		updateDto.Remark = user.Remark
	}

	// 之前的ID怎么是小写
	global.DB.Model(&user).Where("id = ?", updateDto.ID).Updates(map[string]interface{}{
		"nickname": updateDto.Nickname,
		"mail":     updateDto.Mail,
		"phone":    updateDto.Phone,
		"gender":   updateDto.Gender,
		"status":   updateDto.Status,
		"remark":   updateDto.Remark,
	})

	return nil
}

// 删除
func (b *BlogUserService) Delete(deleteDto dto.BlogUserDeleteDto) error {
	var user = &model.BlogUser{}

	err := global.DB.Model(&user).Delete(&user, deleteDto.ID).Error

	return err
}

// 这里该返回什么？
func (b *BlogUserService) List() error {
	var users []model.BlogUser
	err := global.DB.Model(&users).Where("status = ?", 0).Find(&users).Error
	fmt.Println(users)
	return err
}

// QueryByID
func (b *BlogUserService) Query(ID int) error {
	var user []model.BlogUser
	err := global.DB.Model(&user).Where("ID = ?", ID).Find(&user).Error

	return err
}
