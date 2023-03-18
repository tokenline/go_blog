package api

import (
	"go-blog/app/admin/dto"
	"go-blog/common/dto/response"

	"github.com/gin-gonic/gin"
)

type BlogUserApi struct{}

func (BlogUserApi) Query(c *gin.Context) {
	var queryDto dto.BlogUserQueryDto
	c.ShouldBind(&queryDto)

	// TODO :等分页
	// err := userService.Query(queryDto.ID)
	response.Complete(nil, nil, c)
}

func (BlogUserApi) Add(c *gin.Context) {
	var addDto dto.BlogUserAddDto
	c.ShouldBind(&addDto)

	err := userService.Add(addDto)
	response.Complete(nil, err, c)
}

func (BlogUserApi) Update(c *gin.Context) {
	var updateDto dto.BlogUserUpdateDto
	c.ShouldBind(&updateDto)

	err := userService.Update(updateDto)
	response.Complete(nil, err, c)
}

func (BlogUserApi) List(c *gin.Context) {
	vos, err := userService.List()
	response.Complete(vos, err, c)
}

func (BlogUserApi) Delete(c *gin.Context) {
	var deleteDto dto.BlogUserDeleteDto
	c.ShouldBind(&deleteDto)

	err := userService.Delete(deleteDto)
	response.Complete(nil, err, c)
}
