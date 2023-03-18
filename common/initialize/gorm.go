package initialize

import (
	"fmt"
	"go-blog/common/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() *gorm.DB {
	dsn := global.AppConfig.Database.Default
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect to the database")
	}
	fmt.Println("connect successed")
	return db
}
