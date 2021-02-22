package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"ucenter/s/db/connect"
	"ucenter/web/bootstrap"
)

var db = connect.Connect()

func Configure(b *bootstrap.Bootstrapper) {
	mvc.Configure(b.Party("/"), IndexRoute)
	mvc.Configure(b.Party("/user"), UserRoute)
}
