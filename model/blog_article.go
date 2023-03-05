package model

import "gorm.io/gorm"

type BlogArticle struct {
	gorm.Model
	Title     string //文章标题
	Content   string //文章内容
	ViewCount int    //阅读量
	LikeCount int    //点击量
	IsTop     bool   //是否置顶
	Status    int    //状态（默认：0；启用：0；禁用：1）
	Remark    string //备注
}
