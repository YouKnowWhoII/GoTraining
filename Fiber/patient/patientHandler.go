package patient

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func CreatePatient(c *fiber.Ctx) error {
	patient := new(Patient)
	if err := c.BodyParser(patient); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Create(&patient)
	return c.JSON(patient)
}

func GetPatients(c *fiber.Ctx) error {
	var patients []Patient
	db.Find(&patients)
	return c.JSON(patients)
}

func GetPatient(c *fiber.Ctx) error {
	id := c.Params("id")
	var patient Patient
	db.Find(&patient, id)
	return c.JSON(patient)
}

func UpdatePatient(c *fiber.Ctx) error {
	id := c.Params("id")
	var patient Patient
	db.First(&patient, id)
	if patient.ID == 0 {
		return c.Status(404).SendString("No Patient Found with ID")
	}
	updatedPatient := new(Patient)
	if err := c.BodyParser(updatedPatient); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Model(&patient).Updates(updatedPatient)
	return c.JSON(patient)
}

func DeletePatient(c *fiber.Ctx) error {
	id := c.Params("id")
	var patient Patient
	db.First(&patient, id)
	if patient.ID == 0 {
		return c.Status(404).SendString("No Patient Found with ID")
	}
	db.Delete(&patient)
	return c.SendString("Patient Successfully deleted")
}
