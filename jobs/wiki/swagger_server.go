package main

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	_ "ucenter/docs"
)

func main() {
	app := iris.New()

	url := swagger.URL("http://localhost:8090/swagger/doc.json")
	app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler, url))

	// http://localhost:8090/swagger/index.html
	app.Listen(":8090")
}
