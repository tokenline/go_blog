package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResult struct {
	Code    int         //状态码（业务层状态）
	Success bool        //是否成功
	Data    interface{} //数据
	Message string      //消息
}

func Result(code int, success bool, data interface{}, message string, c *gin.Context) {
	//统一返回200
	c.JSON(http.StatusOK, JsonResult{
		Code:    code,
		Success: success,
		Data:    data,
		Message: message,
	})
}

// 请求完成
func Complete(data interface{}, err error, c *gin.Context) {
	if err != nil {
		// 失败
		FailWithDetail(nil, err.Error(), c)
	} else {
		// 成功
		SuccessWithDetail(nil, "请求成功", c)
	}
}

func FailWithDetail(data interface{}, message string, c *gin.Context) {
	Result(200, false, data, message, c)
}

func SuccessWithDetail(data interface{}, message string, c *gin.Context) {
	Result(200, true, data, message, c)
}