package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testRedirect() {
	App.GET("/redirect1", func(c *gin.Context) {
		c.Request.URL.Path = "/redirect2"
		App.HandleContext(c)
	})

	App.GET("/redirect2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
}
