package routes

import (
	"GO/Fiber/handlers"
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

	// Protected route
	protected := app.Group("/download")
	protected.Use(handlers.CheckPermissions)

	app.Get("/download/wifi", limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Hour,
	}), func(c *fiber.Ctx) error {
		return c.Download("Fiber/public/wifi.txt")
	})
}
