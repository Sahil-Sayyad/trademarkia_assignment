package model

import "gorm.io/gorm"

type Inventory struct {
    gorm.Model
    ProductID   uint
    StockLevel  uint
    LowStockThreshold uint
}
