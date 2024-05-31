package models

type User struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password    []byte `json:"-"` // []byte because we will store hashed passwords
	Permissions string `gorm:"type:varchar(100)" json:"permissions"`
}
