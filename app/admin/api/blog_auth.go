package api

type BlogAuthApi struct{}

// @Tags 授权
// @Summary 登录
// @Param data body dto.LoginDto true "请求参数"
// @Success 200 {object} response.JsonResult
// @Router /blog/auth/login [post]
func (*BlogAuthApi) Login() {
	//传入参数

	//service处理

	//返回
}
