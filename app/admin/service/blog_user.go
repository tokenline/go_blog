package service

import (
	"go-blog/app/admin/dto"
	"go-blog/common/dto/request"
	"go-blog/common/global"
	"go-blog/constants"
	"go-blog/model"
)

type BlogUserService struct{}

// QueryByID
func (b *BlogUserService) Query(query dto.BlogUserQuery) (list []dto.BlogUserVo, totalCount int64, err error) {

	db := global.DB.Model(&model.BlogUser{})

	offset := (query.PageNum - 1) * query.PageSize

	//查询数据
	err = db.Offset(offset).Limit(query.PageSize).Find(&list).Error

	if err != nil {
		return nil, 0, err
	}

	//总条数
	db.Count(&totalCount)

	return list, totalCount, err
}

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

// 删除
func (b *BlogUserService) Delete(deleteDto request.IdInfoDto) error {
	var user = &model.BlogUser{}

	err := global.DB.Model(&user).Delete(&user, deleteDto.ID).Error

	return err
}

func (b *BlogUserService) List() (vos []dto.BlogUserVo, err error) {

	err = global.DB.Model(&model.BlogUser{}).Where("status = ?", constants.UserStatusEnable).Find(&vos).Error
	return vos, err
}
