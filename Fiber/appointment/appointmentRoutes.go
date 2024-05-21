package appointment

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.Post("/appointments", CreateAppointment)
	app.Get("/appointments", GetAppointments)
	app.Get("/appointments/:id", GetAppointment)
	app.Put("/appointments/:id", UpdateAppointment)
	app.Delete("/appointments/:id", DeleteAppointment)
}
