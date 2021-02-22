package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"ucenter/repositories"
	"ucenter/services"
	"ucenter/web/controllers"
)

func IndexRoute(app *mvc.Application) {
	repo := repositories.NewUserRepository(db)
	punchService := services.NewUserService(repo)
	app.Register(punchService)
	app.Handle(new(controllers.IndexController))
}
