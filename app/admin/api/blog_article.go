package api

import "github.com/gin-gonic/gin"

type BlogArticleApi struct{}

func (b *BlogArticleApi) Query(c *gin.Context) {}

func (b *BlogArticleApi) Add(c *gin.Context) {}

func (b *BlogArticleApi) Update(c *gin.Context) {}

func (b *BlogArticleApi) Delete(c *gin.Context) {}

func (b *BlogArticleApi) List(c *gin.Context) {}
