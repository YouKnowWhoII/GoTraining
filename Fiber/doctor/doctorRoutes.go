package doctor

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.Post("/doctors", CreateDoctor)
	app.Get("/doctors", GetDoctors)
	app.Get("/doctors/:id", GetDoctor)
	app.Put("/doctors/:id", UpdateDoctor)
	app.Delete("/doctors/:id", DeleteDoctor)
}
