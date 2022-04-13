package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  string `gorm:"not null"`
	IsAdmin   bool
	ID        uint
	FirstName string
	LastName  string
	Phone     string
	Username  string `gorm:"unique"`
}
