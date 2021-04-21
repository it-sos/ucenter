package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"ucenter/repositories"
	"ucenter/services"
	"ucenter/web/controllers"
)

func UserRoute(app *mvc.Application) {
	repo := repositories.NewUserRepository()
	userService := services.NewUserService(repo)
	app.Register(userService)
	app.Handle(new(controllers.UserController))
}
