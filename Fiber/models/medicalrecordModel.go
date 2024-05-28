package models

import (
	"github.com/jinzhu/gorm"
)

type MedicalRecord struct {
	gorm.Model
	PatientID uint
	Patient   Patient `gorm:"foreignKey:PatientID"`
	DoctorID  uint
	Doctor    Doctor `gorm:"foreignKey:DoctorID"`
	Record    string `json:"record"`
}
