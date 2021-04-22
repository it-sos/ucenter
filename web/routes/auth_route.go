package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"ucenter/repositories"
	"ucenter/services"
	"ucenter/web/controllers"
)

func authRoute(app *mvc.Application) {
	repo := repositories.NewUserRepository()
	userService := services.NewUserService(repo)
	authService := services.NewAuthService(userService)
	app.Register(authService)
	app.Handle(new(controllers.UserController))
}
