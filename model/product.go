package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model 
	Name string `gorm:"unique;not null"`
	Description string 
	Price float64
	Quantity uint 
	Orders []Order `gorm:"many2many:order_products;"` 
	Inventory Inventory `gorm:"foreignKey:ProductID"`
}