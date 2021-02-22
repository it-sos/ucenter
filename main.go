package main

import (
	"github.com/kataras/iris/v12"
	"ucenter/web/bootstrap"
	"ucenter/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("punch-in", "peng.yu@yibuyiyin.com")
	app.Bootstrap()
	app.Configure(routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080", iris.WithOptimizations)
}
