package api

import (
	"go-blog/app/admin/dto"
	"go-blog/common/dto/response"

	"github.com/gin-gonic/gin"
)

type BlogAuthApi struct{}

// @Tags 授权
// @Summary 登录
// @Param data body dto.LoginDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /blog/auth/login [post]
func (*BlogAuthApi) Login(c *gin.Context) {
	//传入参数
	var login dto.LoginDto
	c.ShouldBind(&login)

	//service处理
	token, err := authService.Login(login)

	//返回
	response.Complete(token, err, c)
}
