package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloworld() {
	App.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})
}
