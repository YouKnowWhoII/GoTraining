package appointment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type Appointment struct {
	gorm.Model
	PatientID uint   `json:"patient_id"`
	DoctorID  uint   `json:"doctor_id"`
	Date      string `json:"date"`
	Time      string `json:"time"`
}

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func CreateAppointment(c *fiber.Ctx) error {
	appointment := new(Appointment)
	if err := c.BodyParser(appointment); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Create(&appointment)
	return c.JSON(appointment)
}

func GetAppointments(c *fiber.Ctx) error {
	var appointments []Appointment
	db.Find(&appointments)
	return c.JSON(appointments)
}

func GetAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	var appointment Appointment
	db.Find(&appointment, id)
	return c.JSON(appointment)
}

func UpdateAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	var appointment Appointment
	db.First(&appointment, id)
	if appointment.ID == 0 {
		return c.Status(404).SendString("No Appointment Found with ID")
	}
	updatedAppointment := new(Appointment)
	if err := c.BodyParser(updatedAppointment); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Model(&appointment).Updates(updatedAppointment)
	return c.JSON(appointment)
}

func DeleteAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	var appointment Appointment
	db.First(&appointment, id)
	if appointment.ID == 0 {
		return c.Status(404).SendString("No Appointment Found with ID")
	}
	db.Delete(&appointment)
	return c.SendString("Appointment Successfully deleted")
}
