package model

import "gorm.io/gorm"

type BlogCategory struct {
	gorm.Model
	CategoryName string //分类名
	Sort         int    // 排序
	Status       int    // 状态（默认：0；启用：0；禁用：1）
	Remark       string // 备注
}
