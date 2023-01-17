package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email" form:"email" gorm:"unique;not null"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Password  string `json:"password" form:"password" gorm:"not null"`
}
