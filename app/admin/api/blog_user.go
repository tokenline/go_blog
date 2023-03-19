package api

import (
	"go-blog/app/admin/dto"
	"go-blog/common/dto/request"
	"go-blog/common/dto/response"

	"github.com/gin-gonic/gin"
)

type BlogUserApi struct{}

// @Tags 用户
// @Summary 用户查询
// @Param data query dto.BlogUserQuery true "请求参数"
// @Success 200 {object} response.JsonResult{data=response.PageResult{list=[]dto.BlogUserVo}}
// @Router /blog/user/query [get]
func (BlogUserApi) Query(c *gin.Context) {

	var query dto.BlogUserQuery
	c.ShouldBind(&query)

	list, totalCount, err := userService.Query(query)

	response.Complete(response.PageResult{List: list, TotalCount: int(totalCount)}, err, c)
}

// @Tags 用户
// @Summary 增加用户
// @Param data body dto.BlogUserAddDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /blog/user/add [post]
func (BlogUserApi) Add(c *gin.Context) {
	var addDto dto.BlogUserAddDto
	c.ShouldBind(&addDto)

	err := userService.Add(addDto)
	response.Complete(nil, err, c)
}

// @Tags 用户
// @Summary 更新用户信息
// @Param data body dto.BlogUserUpdateDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /blog/user/update [post]
func (BlogUserApi) Update(c *gin.Context) {
	var updateDto dto.BlogUserUpdateDto
	c.ShouldBind(&updateDto)

	err := userService.Update(updateDto)
	response.Complete(nil, err, c)
}

func (BlogUserApi) Detail() {

}

// @Tags 用户
// @Summary 软删除指定用户
// @Param data body request.IdInfoDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /blog/user/delete [post]
func (BlogUserApi) Delete(c *gin.Context) {
	var deleteDto request.IdInfoDto
	c.ShouldBind(&deleteDto)

	err := userService.Delete(deleteDto)
	response.Complete(nil, err, c)
}

// @Tags 用户
// @Summary 返回1个全部状态为启用用户的数组
// @Success 200 {object} response.JsonResult{data=[]dto.BlogUserVo}
// @Router /blog/user/list [get]
func (BlogUserApi) List(c *gin.Context) {
	vos, err := userService.List()
	response.Complete(vos, err, c)
}
