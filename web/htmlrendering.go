package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index() {
	App.LoadHTMLGlob("./templates/*")
	App.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": "admin",
			"pwd":  "123456",
		})
	})
}
