package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"ucenter/web/bootstrap"
)

func Configure(b *bootstrap.Bootstrapper) {
	mvc.Configure(b.Party("/"), indexRoute)
	mvc.Configure(b.Party("/users"), userRoute)
	mvc.Configure(b.Party("/auths"), authRoute)
}
