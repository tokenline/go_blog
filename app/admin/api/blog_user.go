package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogUserApi struct{}

func (BlogUserApi) Query(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "1231231",
	})
}
