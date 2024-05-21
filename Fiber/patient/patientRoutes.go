package patient

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.Post("/patients", CreatePatient)
	app.Get("/patients", GetPatients)
	app.Get("/patients/:id", GetPatient)
	app.Put("/patients/:id", UpdatePatient)
	app.Delete("/patients/:id", DeletePatient)
}
