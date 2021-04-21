package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"ucenter/web/bootstrap"
)

func Configure(b *bootstrap.Bootstrapper) {
	mvc.Configure(b.Party("/"), IndexRoute)
	mvc.Configure(b.Party("/users"), UserRoute)
}
