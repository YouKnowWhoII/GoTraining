package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterDownloadRoutes(app *fiber.App) {
	app.Get("/download/manual.txt", func(c *fiber.Ctx) error {
		return c.Download("./public/manual.pdf")
	})
}
