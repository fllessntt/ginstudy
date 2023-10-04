package web

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func serverPush() {
	App.Static("/assets", "./assets")
	App.SetHTMLTemplate(html)
	App.GET("/push", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			//使用 pusher.Push()做服务推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Faild to push:%v\n", err)
				//return
			}
		}
		c.HTML(200, "https", gin.H{"status": "success"})
	})
}
