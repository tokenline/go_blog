package api

import (
	"go-blog/app/admin/dto"
	"go-blog/common/dto/response"

	"github.com/gin-gonic/gin"
)

type BlogUserApi struct{}

func (BlogUserApi) Query(c *gin.Context) {

	var query dto.BlogUserQuery
	c.ShouldBind(&query)

	list, totalCount, err := userService.Query(query)

	response.Complete(response.PageResult{List: list, TotalCount: int(totalCount)}, err, c)
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

func (BlogUserApi) Detail() {

}

func (BlogUserApi) Delete(c *gin.Context) {
	var deleteDto dto.BlogUserDeleteDto
	c.ShouldBind(&deleteDto)

	err := userService.Delete(deleteDto)
	response.Complete(nil, err, c)
}

func (BlogUserApi) List(c *gin.Context) {
	vos, err := userService.List()
	response.Complete(vos, err, c)
}
