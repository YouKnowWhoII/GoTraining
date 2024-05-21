package main

import (
	"GO/Fiber/appointment"
	"GO/Fiber/doctor"
	"GO/Fiber/medicalrecord"
	"GO/Fiber/patient"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&patient.Patient{}, &doctor.Doctor{}, &appointment.Appointment{}, &medicalrecord.MedicalRecord{})
	patient.SetDB(db)
	doctor.SetDB(db)
	appointment.SetDB(db)
	medicalrecord.SetDB(db)
}

func main() {
	app := fiber.New()

	// Patient routes
	app.Post("/patients", patient.CreatePatient)
	app.Get("/patients", patient.GetPatients)
	app.Get("/patients/:id", patient.GetPatient)
	app.Put("/patients/:id", patient.UpdatePatient)
	app.Delete("/patients/:id", patient.DeletePatient)

	// Doctor routes
	app.Post("/doctors", doctor.CreateDoctor)
	app.Get("/doctors", doctor.GetDoctors)
	app.Get("/doctors/:id", doctor.GetDoctor)
	app.Put("/doctors/:id", doctor.UpdateDoctor)
	app.Delete("/doctors/:id", doctor.DeleteDoctor)

	// Appointment routes
	app.Post("/appointments", appointment.CreateAppointment)
	app.Get("/appointments", appointment.GetAppointments)
	app.Get("/appointments/:id", appointment.GetAppointment)
	app.Put("/appointments/:id", appointment.UpdateAppointment)
	app.Delete("/appointments/:id", appointment.DeleteAppointment)

	// MedicalRecord routes
	app.Post("/medicalrecords", medicalrecord.CreateMedicalRecord)
	app.Get("/medicalrecords", medicalrecord.GetMedicalRecords)
	app.Get("/medicalrecords/:id", medicalrecord.GetMedicalRecord)
	app.Put("/medicalrecords/:id", medicalrecord.UpdateMedicalRecord)
	app.Delete("/medicalrecords/:id", medicalrecord.DeleteMedicalRecord)

	app.Listen(":3000")
}
