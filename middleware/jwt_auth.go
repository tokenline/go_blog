package middleware

import (
	"fmt"
	"go-blog/common/dto/response"
	"go-blog/common/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 授权
func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.Request.Header.Get("Authorization")
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", -1)

		if tokenStr == "" {
			response.FailWithCode(http.StatusUnauthorized, "未授权", ctx)
			ctx.Abort() //结束后续操作
			return
		}

		token, err := utils.Jwt.ParseToken(tokenStr)
		if err != nil {
			response.FailWithCode(http.StatusUnauthorized, "授权失败，错误token", ctx)
			ctx.Abort() //结束后续操作
			return
		}

		// 获取token信息
		claim, ok := token.Claims.(*utils.UserAuthClaims)
		if !ok {
			response.FailWithCode(http.StatusUnauthorized, "授权失败，解析失败", ctx)
			ctx.Abort() //结束后续操作
			return
		}
		fmt.Println(claim)

		ctx.Next() //接着下一步
	}
}
