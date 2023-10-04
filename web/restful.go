package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const student = "/student"

func restful() {
	App.GET(student, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "查询学生信息成功"})
	})
	App.POST(student, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "创建学生信息成功"})
	})
	App.PUT(student, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "更新学生信息成功"})
	})
	App.DELETE(student, func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{"message": "删除学生信息成功"})
	})
}
