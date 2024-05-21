package appointment

import "github.com/jinzhu/gorm"

type Appointment struct {
	gorm.Model
	PatientID uint   `json:"patient_id"`
	DoctorID  uint   `json:"doctor_id"`
	Date      string `json:"date"`
	Time      string `json:"time"`
}
