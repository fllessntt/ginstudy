package web

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Userinfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func (u *Userinfo) IsEmpty() bool {
	return u.Username == "" || u.Password == ""
}

func testBind() {
	App.Any("/user", func(c *gin.Context) {
		var u Userinfo
		//ShouldBind()
		//如果是GET请求，只使用Form绑定引擎（Query）
		//如果是POST请求，首先检查content-type是否为JSON或XML，然后再使用Form（form-data）
		err := c.ShouldBind(&u)
		//err := c.ShouldBindQuery(&u)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		} else {
			if u.IsEmpty() {
				c.JSON(http.StatusOK, gin.H{"info": "用户名或密码为空"})
			} else {
				c.JSON(http.StatusOK, u)
			}
		}
		log.Printf("%#v\n", u)
	})
}
