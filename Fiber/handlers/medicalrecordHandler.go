package handlers

import (
	"GO/Fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func SetMedicalRecordDB(database *gorm.DB) {
	db = database
}

func CreateMedicalRecord(c *fiber.Ctx) error {
	medicalRecord := new(models.MedicalRecord)
	if err := c.BodyParser(medicalRecord); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Find the patient with the provided ID
	var patient models.Patient
	db.First(&patient, medicalRecord.PatientID)
	if patient.ID == 0 {
		return c.Status(404).SendString("No Patient Found with ID")
	}

	// Find the doctor with the provided ID
	var doctor models.Doctor
	db.First(&doctor, medicalRecord.DoctorID)
	if doctor.ID == 0 {
		return c.Status(404).SendString("No Doctor Found with ID")
	}

	// Associate the medical record with the existing patient and doctor
	medicalRecord.Patient = patient
	medicalRecord.Doctor = doctor

	db.Create(&medicalRecord)
	return c.JSON(medicalRecord)
}

func GetMedicalRecords(c *fiber.Ctx) error {
	var medicalRecords []models.MedicalRecord
	db.Preload("Patient").Preload("Doctor").Find(&medicalRecords)
	return c.JSON(medicalRecords)
}

func GetMedicalRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var medicalRecord models.MedicalRecord
	db.Preload("Patient").Preload("Doctor").Find(&medicalRecord, id)
	return c.JSON(medicalRecord)
}

func UpdateMedicalRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var medicalRecord models.MedicalRecord
	db.First(&medicalRecord, id)
	if medicalRecord.ID == 0 {
		return c.Status(404).SendString("No MedicalRecord Found with ID")
	}
	updatedMedicalRecord := new(models.MedicalRecord)
	if err := c.BodyParser(updatedMedicalRecord); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Find the patient with the provided ID
	var patient models.Patient
	db.First(&patient, updatedMedicalRecord.PatientID)
	if patient.ID == 0 {
		return c.Status(404).SendString("No Patient Found with ID")
	}

	// Find the doctor with the provided ID
	var doctor models.Doctor
	db.First(&doctor, updatedMedicalRecord.DoctorID)
	if doctor.ID == 0 {
		return c.Status(404).SendString("No Doctor Found with ID")
	}

	// Associate the updated medical record with the existing patient and doctor
	updatedMedicalRecord.Patient = patient
	updatedMedicalRecord.Doctor = doctor

	db.Model(&medicalRecord).Updates(updatedMedicalRecord)
	return c.JSON(medicalRecord)
}

func DeleteMedicalRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var medicalRecord models.MedicalRecord
	db.First(&medicalRecord, id)
	if medicalRecord.ID == 0 {
		return c.Status(404).SendString("No MedicalRecord Found with ID")
	}
	db.Delete(&medicalRecord)
	return c.SendString("MedicalRecord Successfully deleted")
}
