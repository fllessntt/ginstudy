package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testQuery() {
	App.GET("/query", func(c *gin.Context) {
		name := c.Query("name")
		pwd := c.Query("pwd")
		//log.Printf("name:%s ; pwd:%s\n", name, pwd)
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"pwd":  pwd,
		})
	})
	App.POST("/index", func(c *gin.Context) {
		name := c.PostForm("name")
		pwd := c.PostForm("pwd")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": name,
			"pwd":  pwd,
		})
	})
}
