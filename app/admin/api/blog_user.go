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

	err := userService.Query(queryDto.ID)
	response.Complete(nil, err, c)
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

// List这个API应该不用传入上下文？
// 我尝试删除 *gin.Context结果rounter_admin.go中对应的调用List
func (BlogUserApi) List(c *gin.Context) {
	err := userService.List()
	response.Complete(nil, err, nil)
}

func (BlogUserApi) Delete(c *gin.Context) {
	var deleteDto dto.BlogUserDeleteDto
	c.ShouldBind(&deleteDto)

	err := userService.Delete(deleteDto)
	response.Complete(nil, err, c)
}
