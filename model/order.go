package model

import "gorm.io/gorm"

type Order struct {
    gorm.Model
    UserID      uint  `gorm:"foreignKey:UserID;references:ID"`
    User        User  `gorm:"foreignKey:UserID;references:ID"`
    Products    []Product `gorm:"many2many:order_products;"`
    TotalPrice  float64
    Status      string   `gorm:"default:'pending'"` // Default status is 'pending'
}
