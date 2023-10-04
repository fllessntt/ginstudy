package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func anyAndNo() {
	App.Any("/any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "一切皆有可能"})
	})
	App.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "404"})
	})
}

const person = "/person"
const actor = "/actor"

func group() {
	personRouter := App.Group(person)
	personRouter.GET("/query", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "获取人员信息成功"})
	})
	personRouter.GET("/save", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "创建人员信息成功"})
	})
	personRouter.GET("/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "更新人员信息成功"})
	})
	personRouter.GET("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "删除人员信息成功"})
	})
	actorRouter := personRouter.Group(actor)
	actorRouter.GET("/love", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "最爱的演员是丁勇岱"})
	})
}
