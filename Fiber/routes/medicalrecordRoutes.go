package routes

import (
	"GO/Fiber/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func RegisterMedicalRecordRoutes(app *fiber.App) {
	app.Post("/medicalrecords", handlers.CreateMedicalRecord)
	app.Get("/medicalrecords", handlers.GetMedicalRecords)
	app.Get("/medicalrecords/:id", handlers.GetMedicalRecord)
	//route to return a text file with a medical record by id
	app.Get("/medicalrecords/:id/download", limiter.New(limiter.Config{
		Max:        5, // limit each IP to 5 requests per minute
		Expiration: 1 * time.Minute,
	}), handlers.GetMedicalRecordFile)
	app.Put("/medicalrecords/:id", handlers.UpdateMedicalRecord)
	app.Delete("/medicalrecords/:id", handlers.DeleteMedicalRecord)
}
