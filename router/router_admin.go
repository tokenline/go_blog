package router

import (
	"go-blog/app/admin/api"

	"github.com/gin-gonic/gin"
)

func InitAdminRouters(routerGroup *gin.RouterGroup) {
	var blogUserApi = api.BlogUserApi{}
	blogUser := routerGroup.Group("/blog/user")
	{
		blogUser.GET("query", blogUserApi.Query)
		blogUser.POST("add", blogUserApi.Add)
		blogUser.POST("update", blogUserApi.Query)
		blogUser.POST("delete", blogUserApi.Query)
		blogUser.GET("detail", blogUserApi.Query)
		blogUser.GET("list", blogUserApi.Query)
	}

}
