package api

import "github.com/gin-gonic/gin"

type BlogCategoryApi struct{}

func (b *BlogCategoryApi) Query(c *gin.Context) {}

func (b *BlogCategoryApi) Add(c *gin.Context) {}

func (b *BlogCategoryApi) Update(c *gin.Context) {}

func (b *BlogCategoryApi) Delete(c *gin.Context) {}

func (b *BlogCategoryApi) List(c *gin.Context) {}
