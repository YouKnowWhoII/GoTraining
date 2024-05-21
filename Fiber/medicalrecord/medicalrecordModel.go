package medicalrecord

import (
	"GO/Fiber/doctor"
	"GO/Fiber/patient"
	"github.com/jinzhu/gorm"
)

type MedicalRecord struct {
	gorm.Model
	PatientID uint
	Patient   patient.Patient `gorm:"foreignKey:PatientID"`
	DoctorID  uint
	Doctor    doctor.Doctor `gorm:"foreignKey:DoctorID"`
	Record    string        `json:"record"`
}
