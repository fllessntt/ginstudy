package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func calculateTimeCost(c *gin.Context) {
	log.Println("calculateTimeCost...")
	start := time.Now()
	//c.Abort() //阻止调用后续的处理函数
	c.Next() //调用后续的处理函数
	cost := time.Since(start)
	log.Printf("cost:%v\n", cost)
}

func ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func testHook() {
	App.GET("/cost", calculateTimeCost, ok)
}
