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
		blogUser.POST("update", blogUserApi.Update)
		blogUser.POST("delete", blogUserApi.Delete)
		//blogUser.GET("detail", blogUserApi.Query)
		blogUser.GET("list", blogUserApi.List)
	}

	var blogArticleApi = api.BlogArticleApi{}
	blogArticle := routerGroup.Group("/blog/article")
	{
		blogArticle.GET("query", blogArticleApi.Query)
		blogArticle.POST("add", blogArticleApi.Add)
		blogArticle.POST("update", blogArticleApi.Update)
		blogArticle.POST("delete", blogArticleApi.Delete)
		//blogArticle.GET("detail", blogArticleApi.Query) 这里需要吗？
		blogArticle.GET("list", blogArticleApi.List)
	}

	var blogCategoryApi = api.BlogCategoryApi{}
	blogCategory := routerGroup.Group("/blog/category")
	{
		blogCategory.GET("query", blogCategoryApi.Query)
		blogCategory.POST("add", blogCategoryApi.Add)
		blogCategory.POST("update", blogCategoryApi.Update)
		blogCategory.POST("delete", blogCategoryApi.Delete)
		//blogCategory.GET("detail", blogCategoryApi.Query) 这里需要吗？
		blogCategory.GET("list", blogCategoryApi.List)
	}

	var blogAuthApi = api.BlogAuthApi{}
	blogAuth := routerGroup.Group("/blog/auth")
	{
		blogAuth.POST("login", blogAuthApi.Login)
	}

}
