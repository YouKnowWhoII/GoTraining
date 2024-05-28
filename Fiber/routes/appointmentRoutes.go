package routes

import (
	"GO/Fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterAppointmentRoutes(app *fiber.App) {
	app.Post("/appointments", handlers.CreateAppointment)
	app.Get("/appointments", handlers.GetAppointments)
	app.Get("/appointments/:id", handlers.GetAppointment)
	app.Put("/appointments/:id", handlers.UpdateAppointment)
	app.Delete("/appointments/:id", handlers.DeleteAppointment)
}
