package routes

import (
	"GO/Fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterPatientRoutes(app *fiber.App) {
	app.Post("/patients", handlers.CreatePatient)
	app.Get("/patients", handlers.GetPatients)
	app.Get("/patients/:id", handlers.GetPatient)
	app.Put("/patients/:id", handlers.UpdatePatient)
	app.Delete("/patients/:id", handlers.DeletePatient)
}
