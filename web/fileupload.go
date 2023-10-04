package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func uploadFile() {
	//文件上传测试页面
	App.GET("/upload.html", func(c *gin.Context) {
		info, _ := c.Get("info")
		multiinfo, _ := c.Get("info")
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"info":      info,
			"multiinfo": multiinfo,
		})
	})
	//单文件上传
	App.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		log.Println(file.Filename)
		dst := fmt.Sprintf("/Users/yijiuchow/Desktop/%s", file.Filename)
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.Set("info", fmt.Sprintf("%s uploaded success!", file.Filename))
		c.Redirect(http.StatusMovedPermanently, "upload.html")
	})
	//多文件上传
	App.POST("/multiupload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		files := form.File["file"]
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("/Users/yijiuchow/Desktop/%d_%s", index, file.Filename)
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}
		c.Set("multiinfo", fmt.Sprintf("%d files uploaded success!", len(files)))
		c.Redirect(http.StatusMovedPermanently, "upload.html")
	})
}
