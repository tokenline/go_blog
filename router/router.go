package router

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	routers := gin.Default()

	InitAdminRouters(&routers.RouterGroup)
	return routers
}
