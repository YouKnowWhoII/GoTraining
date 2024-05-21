package medicalrecord

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.Post("/medicalrecords", CreateMedicalRecord)
	app.Get("/medicalrecords", GetMedicalRecords)
	app.Get("/medicalrecords/:id", GetMedicalRecord)
	app.Put("/medicalrecords/:id", UpdateMedicalRecord)
	app.Delete("/medicalrecords/:id", DeleteMedicalRecord)
}
