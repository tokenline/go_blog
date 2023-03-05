package model

import "gorm.io/gorm"

type BlogArticleCategory struct {
	gorm.Model
	ArticleId  int // 文章ID
	CategoryID int // 分类ID
}
