package routes

import (
	"GO/Fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App) {

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Get("/user", handlers.User)
	app.Post("/logout", handlers.Logout)

}
