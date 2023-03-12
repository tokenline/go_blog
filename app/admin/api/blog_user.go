package api

import (
	"errors"
	"go-blog/app/admin/dto"
	"go-blog/common/dto/response"

	"github.com/gin-gonic/gin"
)

type BlogUserApi struct{}

func (*BlogUserApi) Query(c *gin.Context) {

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "1231231",
	// })
	response.Complete(nil, errors.New("错误了"), c)
}

// 添加用户
func (*BlogUserApi) Add(c *gin.Context) {
	// 传入参数
	var addDto dto.BlogUserAddDto
	// 请求参数赋值到结构体中
	c.ShouldBindJSON(&addDto)

	// 业务处理（service）
	err := userService.Add(addDto)

	// 返回结果
	response.Complete(nil, err, c)
}

func (BlogUserApi) Query(c *gin.Context)
