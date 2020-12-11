package main

import (
	"ucenter/web/bootstrap"
	"ucenter/web/middleware/identity"
	"ucenter/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("UCenter", "peng.yu@yibuyiyin.com")
	app.Bootstrap()
	app.Configure(routes.Configure, identity.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}