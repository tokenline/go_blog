package main

import (
	"go-blog/common/global"
	"go-blog/common/initialize"
	"go-blog/router"
)

func main() {
	//初始化配置 yaml
	global.AppConfig = initialize.InitAppConfig()

	//初始化数据库
	global.DB = initialize.InitDB()

	//初始化路由
	var router = router.InitRouters()

	router.Run()
}
