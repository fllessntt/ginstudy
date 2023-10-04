package web

import "github.com/gin-gonic/gin"

var (
	App = gin.Default()
)

func RunWebApplication() {
	helloworld()
	restful()
	index()
	testQuery()
	anyAndNo()
	group()
	testHook()
	testGoRoutine()
	testBind()
	uploadFile()
	testRedirect()
	checkbox()
	variousJson()
	serverPush()
	variousSerialize()
	_ = App.Run("0.0.0.0:7777")
	//_ = App.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
