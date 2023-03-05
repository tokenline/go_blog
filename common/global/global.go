package global

import (
	"go-blog/config"

	"gorm.io/gorm"
)

var (
	AppConfig *config.AppConfig //配置
	DB        *gorm.DB          //数据库连接
)
