package utils

import (
	"github.com/Sahil-Sayyad/Trademarkia/database"
	"github.com/Sahil-Sayyad/Trademarkia/model"
)

// FindProductByID retrieves a product by its ID
func FindProductByID(id uint) (*model.Product, error) {

	var product model.Product

	if err := database.DB.First(&product, id).Error; 
	err != nil {

		return nil, err
	}
	
	return &product, nil
}
