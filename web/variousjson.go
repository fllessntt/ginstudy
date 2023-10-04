package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func variousJson() {
	App.GET("/ascii", func(c *gin.Context) {
		c.AsciiJSON(200, gin.H{"lang": "语言", "country": "国家"})
	})
	App.GET("/pure", func(c *gin.Context) {
		c.PureJSON(200, gin.H{"html": "<b>Hello,World</b>"})
	})
	App.GET("/common", func(c *gin.Context) {
		c.JSON(200, gin.H{"html": "<b>Hello,World</b>"})
	})
	App.GET("/secure", func(c *gin.Context) {
		names := []string{"dyd", "zrf"}
		//输出 while(1);["dyd","zrf"]
		c.SecureJSON(200, names)
	})
}

// 输出 while(1);["dyd","zrf"]
func variousSerialize() {
	App.GET("/struct", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Dyd"
		msg.Message = "hey"
		msg.Number = 123
		c.JSON(200, msg)
	})

	App.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "hey", "status": 200})
	})

	App.GET("/testYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"message": "hey", "status": 200})
	})

	App.GET("protobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在testdata/proto example文件中
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意, 数据在响应中变为二进制数据
		// 将输出被 proto example.Test protobuf 序列化的数据
		c.ProtoBuf(200, data)

	})
}
