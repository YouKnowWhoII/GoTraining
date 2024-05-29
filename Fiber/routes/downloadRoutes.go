package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func RegisterDownloadRoutes(app *fiber.App) {
	app.Get("/download/manual.txt", limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Hour,
	}), func(c *fiber.Ctx) error {
		return c.Download("./public/manual.pdf")
	})
}
