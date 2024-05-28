package handlers

import (
	"GO/Fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func SetDoctorDB(database *gorm.DB) {
	db = database
}

func CreateDoctor(c *fiber.Ctx) error {
	doctor := new(models.Doctor)
	if err := c.BodyParser(doctor); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Create(&doctor)
	return c.JSON(doctor)
}

func GetDoctors(c *fiber.Ctx) error {
	var doctors []models.Doctor
	db.Find(&doctors)
	return c.JSON(doctors)
}

func GetDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	var doctor models.Doctor
	db.Find(&doctor, id)
	return c.JSON(doctor)
}

func UpdateDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	var doctor models.Doctor
	db.First(&doctor, id)
	if doctor.ID == 0 {
		return c.Status(404).SendString("No Doctor Found with ID")
	}
	updatedDoctor := new(models.Doctor)
	if err := c.BodyParser(updatedDoctor); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Model(&doctor).Updates(updatedDoctor)
	return c.JSON(doctor)
}

func DeleteDoctor(c *fiber.Ctx) error {
	id := c.Params("id")
	var doctor models.Doctor
	db.First(&doctor, id)
	if doctor.ID == 0 {
		return c.Status(404).SendString("No Doctor Found with ID")
	}
	db.Delete(&doctor)
	return c.SendString("Doctor Successfully deleted")
}
