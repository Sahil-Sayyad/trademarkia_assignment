package model

/*
	1.User
	2.Admin
	3.Product
	4.Order
*/
import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name     string
	Orders   []Order
}

type Admin struct {
	gorm.Model

	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name     string
}

type Product struct {
	gorm.Model

	Name             string  `gorm:"not null"`
	Price            float64 `gorm:"not null"`
	ShoppingCategory string  `gorm:"not null"`
	Quantity         int     `gorm:"not null"`
}

type Order struct {
	gorm.Model

	UserID     uint
	User       User      `gorm:"foreignKey:UserID"`
	Products   []Product `gorm:"many2many:order_products;"`
	TotalPrice float64
}
