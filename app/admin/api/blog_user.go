package api

import (
	"fmt"
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

func (BlogUserApi) Query(c *gin.Context)
