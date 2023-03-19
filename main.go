package main

import (
	"fmt"
	"go-blog/common/global"
	"go-blog/common/initialize"
	"go-blog/docs"
	"go-blog/router"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//初始化配置 yaml
	global.AppConfig = initialize.InitAppConfig()

	//初始化数据库
	global.DB = initialize.InitDB()

	//初始化路由
	var router = router.InitRouters()

	InitSwagger(router)

	// router.Run()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Printf("Swagger文档：http://127.0.0.1%s/swagger/index.html", server.Addr)
	server.ListenAndServe()
}

func InitSwagger(router *gin.Engine) {
	docs.SwaggerInfo.Title = "go-blog"
	docs.SwaggerInfo.Description = "Swagger Admin API"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
