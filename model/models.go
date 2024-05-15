package model

import (
	"gorm.io/gorm"
)

type User struct {
	
	gorm.Model
	 
	Email 	string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name string 
	//Orders []Order `gorm:"foreignKey:userID"`// One-to-many relationship
}

type Admin struct {
	gorm.Model

	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name string 
}

type Product struct {
	gorm.Model

	Name string `gorm:"not null"`
	Price float64 `gorm:"not null"`
	ShoppingCategory string `gorm:"not null"`
}

type Order struct {
	gorm.Model
	UserID uint
	ProductID uint
	Quantity int 
	TotalCost float64
}