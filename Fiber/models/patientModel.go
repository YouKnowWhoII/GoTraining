package models

import (
	"github.com/jinzhu/gorm"
)

type Patient struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `gorm:"type:varchar(100);unique_index" json:"email"`
	DOB   string `json:"dob"`
}
