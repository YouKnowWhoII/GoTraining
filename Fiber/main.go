package main

import (
	"GO/Fiber/handlers"
	"GO/Fiber/models"
	"GO/Fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//func resetDatabase(db *gorm.DB) {
//	db.DropTableIfExists(&patient.Patient{}, &doctor.Doctor{}, &appointment.Appointment{}, &medicalrecord.MedicalRecord{})
//	db.AutoMigrate(&patient.Patient{}, &doctor.Doctor{}, &appointment.Appointment{}, &medicalrecord.MedicalRecord{})
//}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Patient{}, &models.Doctor{}, &models.Appointment{}, &models.MedicalRecord{}, &models.User{})
	handlers.SetPatientDB(db)
	handlers.SetDoctorDB(db)
	handlers.SetMedicalRecordDB(db)
	handlers.SetAppointmentDB(db)
	handlers.SetAuthDB(db)
}

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	//app.Use(cors.New(cors.Config{
	//	AllowCredentials: true,
	//}))

	//resetDatabase(db)

	routes.RegisterPatientRoutes(app)
	routes.RegisterDoctorRoutes(app)
	routes.RegisterMedicalRecordRoutes(app)
	routes.RegisterAppointmentRoutes(app)
	routes.RegisterDownloadRoutes(app)
	routes.RegisterAuthRoutes(app)

	app.Listen(":3000")
}
