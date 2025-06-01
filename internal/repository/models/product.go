package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Group string
	Price string
}
