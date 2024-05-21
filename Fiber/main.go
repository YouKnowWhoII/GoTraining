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
	db.AutoMigrate(&patient.Patient{}, &doctor.Doctor{}, &appointment.Appointment{}, &medicalrecord.MedicalRecord{})
	patient.SetDB(db)
	doctor.SetDB(db)
	appointment.SetDB(db)
	medicalrecord.SetDB(db)
}

func main() {
	app := fiber.New()

	//resetDatabase(db)

	patient.RegisterRoutes(app)
	doctor.RegisterRoutes(app)
	appointment.RegisterRoutes(app)
	medicalrecord.RegisterRoutes(app)

	app.Listen(":3000")
}
