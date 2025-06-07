package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
	Name     string
	Age      int
	Role     string
}
