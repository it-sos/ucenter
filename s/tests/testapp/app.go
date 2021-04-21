package testapp

import (
	"github.com/kataras/iris/v12"
	"ucenter/web/bootstrap"
)

func App(cfg bootstrap.Configurator) *iris.Application {
	app := bootstrap.New("ucenter", "peng.yu@yibuyiyin.com")
	app.Bootstrap()
	app.Configure(cfg)
	return app.Application
}
