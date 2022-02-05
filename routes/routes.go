package routes

import (
	"golang-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	group := app.Group("/api")
	group.Post("/register", controllers.Register)
	group.Post("/login", controllers.Login)
	group.Get("/user", controllers.User)
	group.Post("/logout", controllers.Logout)
}
