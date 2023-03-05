package model

import "gorm.io/gorm"

type BlogArticleCategory struct {
	gorm.Model
	ArticleId  int // 文章ID
	CategoryId int // 分类ID
}
