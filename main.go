package main

import (
	"github.com/kataras/iris/v12"
	"ucenter/web/bootstrap"
	"ucenter/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("ucenter", "peng.yu@yibuyiyin.com")
	app.Bootstrap()
	app.Configure(routes.Configure)
	return app
}

// @title ucenter api
// @version 1.0
// @description 用户中心接口
// @termsOfService http://www.yibuyiyin.com/terms/

// @contact.name API Support
// @contact.url http://www.yibuyiyin.com/support
// @contact.email peng.yu@yibuyiyin.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	app := newApp()
	app.Listen(":8080", iris.WithOptimizations)
}
