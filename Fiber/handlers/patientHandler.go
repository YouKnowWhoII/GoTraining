package handlers

import (
	"GO/Fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetPatientDB(database *gorm.DB) {
	db = database
}

func CreatePatient(c *fiber.Ctx) error {
	patient := new(models.Patient)
	if err := c.BodyParser(patient); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Create(&patient)
	return c.JSON(patient)
}

func GetPatients(c *fiber.Ctx) error {
	var patients []models.Patient
	db.Find(&patients)
	return c.JSON(patients)
}

func GetPatient(c *fiber.Ctx) error {
	id := c.Params("id")
	var patient models.Patient
	db.Find(&patient, id)
	return c.JSON(patient)
}

func UpdatePatient(c *fiber.Ctx) error {
	id := c.Params("id")
	var patient models.Patient
	db.First(&patient, id)
	if patient.ID == 0 {
		return c.Status(404).SendString("No Patient Found with ID")
	}
	updatedPatient := new(models.Patient)
	if err := c.BodyParser(updatedPatient); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Model(&patient).Updates(updatedPatient)
	return c.JSON(patient)
}

func DeletePatient(c *fiber.Ctx) error {
	id := c.Params("id")
	var patient models.Patient
	db.First(&patient, id)
	if patient.ID == 0 {
		return c.Status(404).SendString("No Patient Found with ID")
	}
	db.Delete(&patient)
	return c.SendString("Patient Successfully deleted")
}
