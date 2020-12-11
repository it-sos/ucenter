package routes

import (
	"ucenter/web/bootstrap"
	"ucenter/web/controllers"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	b.Get("/login", controllers.Login)
}
