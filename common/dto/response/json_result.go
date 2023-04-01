package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResult struct {
	Code    int         `json:"code"`    //状态码（业务层状态）
	Success bool        `json:"success"` //是否成功
	Data    interface{} `json:"data"`    //数据
	Message string      `json:"message"` //消息
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
		FailWithDetail(data, err.Error(), c)
	} else {
		// 成功
		SuccessWithDetail(data, "请求成功", c)
	}
}

func FailWithDetail(data interface{}, message string, c *gin.Context) {
	Result(200, false, data, message, c)
}

func FailWithCode(code int, message string, c *gin.Context) {
	Result(code, false, nil, message, c)
}

func SuccessWithDetail(data interface{}, message string, c *gin.Context) {
	Result(200, true, data, message, c)
}
