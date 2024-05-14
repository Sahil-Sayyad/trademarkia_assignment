package model

import "gorm.io/gorm"

type Order struct {
    gorm.Model
    UserID      uint
    User        User      // Belongs to User
    Products    []Product `gorm:"many2many:order_products;"`
    TotalPrice  float64
    Status      string   `gorm:"default:'pending'"` // Default status is 'pending'
}
