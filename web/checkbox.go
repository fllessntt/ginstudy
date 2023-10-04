package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func checkbox() {
	App.GET("/checkbox", func(c *gin.Context) {
		c.HTML(200, "checkbox.html", nil)
	})
	App.POST("/checkbox", func(c *gin.Context) {
		var form myForm
		err := c.Bind(&form)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"color": form.Colors})
	})
}
