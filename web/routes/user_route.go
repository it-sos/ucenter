package routes

import (
	"github.com/kataras/iris/v12/_examples/mvc/login/web/controllers"
	"github.com/kataras/iris/v12/mvc"
	"ucenter/repositories"
	"ucenter/services"
)

func UserRoute(app *mvc.Application) {
	repo := repositories.NewUserRepository(db)
	userService := services.NewUserService(repo)
	app.Register(userService)
	app.Handle(new(controllers.UserController))
}
