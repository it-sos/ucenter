package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"os"
	"ucenter/s/db/connect"
	"ucenter/web/bootstrap"
)

var db = connect.Connect()

func Configure(b *bootstrap.Bootstrapper) {
	mvc.Configure(b.Party("/"), IndexRoute)
	mvc.Configure(b.Party("/users"), UserRoute)
}

func testApp() *iris.Application {
	os.Chdir("/data1/htdocs/ucenter")
	app := bootstrap.New("ucenter", "peng.yu@yibuyiyin.com")
	app.Bootstrap()
	app.Configure(Configure)
	return app.Application
}
