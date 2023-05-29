package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	Barcode    string
	Price      uint
	CategoryID uint
	Category   Category
}
