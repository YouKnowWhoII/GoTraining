package routes

import (
	"GO/Fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterMedicalRecordRoutes(app *fiber.App) {
	app.Post("/medicalrecords", handlers.CreateMedicalRecord)
	app.Get("/medicalrecords", handlers.GetMedicalRecords)
	app.Get("/medicalrecords/:id", handlers.GetMedicalRecord)
	//route to return a text file with a medical record by id
	app.Get("/medicalrecords/:id/download", handlers.GetMedicalRecordFile)
	app.Put("/medicalrecords/:id", handlers.UpdateMedicalRecord)
	app.Delete("/medicalrecords/:id", handlers.DeleteMedicalRecord)
}
