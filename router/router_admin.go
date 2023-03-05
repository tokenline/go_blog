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
	}

}
