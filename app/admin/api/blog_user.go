package api

import (
	"fmt"
	"go-blog/app/admin/dto"
	"go-blog/common/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogUserApi struct{}

func (BlogUserApi) Query(c *gin.Context) {
	fmt.Println("kaishi yewuchu")
	c.JSON(http.StatusOK, gin.H{
		"message": "1231231",
	})
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
