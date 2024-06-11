package routes

import (
	"GO/Fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterDoctorRoutes(app *fiber.App) {

	app.Get("/doctors", handlers.GetDoctors)
	app.Get("/doctors/:id", handlers.GetDoctor)

	// Protected route
	protected := app.Group("/doctors")
	protected.Use(handlers.CheckPermissions)

	app.Post("/doctors", handlers.CreateDoctor)
	app.Put("/doctors/:id", handlers.UpdateDoctor)
	app.Delete("/doctors/:id", handlers.DeleteDoctor)
}
