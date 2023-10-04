package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func testGoRoutine() {
	App.GET("/async", func(c *gin.Context) {
		tmp := c.Copy()
		// 创建再goroutine中使用的副本
		go func() {
			//模拟长任务
			time.Sleep(5 * time.Second)
			//请注意你使用的是复制的上下文 "tmp", 这是上下文的一个copy
			log.Println("Done! in path " + tmp.Request.URL.Path)

		}()
		c.JSON(http.StatusOK, gin.H{"message": "Done! in path " + c.Request.URL.Path})
	})
	App.GET("/sync", func(c *gin.Context) {
		//模拟长任务
		time.Sleep(5 * time.Second)
		//因为没有使用goroutine, 不需要拷贝上下文
		log.Println("Done! in path " + c.Request.URL.Path)
		c.JSON(http.StatusOK, gin.H{"message": "Done! in path " + c.Request.URL.Path})
	})
}
