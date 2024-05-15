package controller

import (
	"strconv"

	"github.com/Sahil-Sayyad/Trademarkia/database"
	"github.com/Sahil-Sayyad/Trademarkia/model"
	"github.com/Sahil-Sayyad/Trademarkia/utils"
	"github.com/gofiber/fiber/v2"
)

// CreateProduct creates an product
func CreateProduct(c *fiber.Ctx) error {

	var product model.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	//Create product in the database
	result := database.DB.Create(&product)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create product"})
	}

	return c.JSON(product)
}

// UpdateProduct updates an existing product
func UpdateProduct(c *fiber.Ctx) error {

	//Parse product ID from URL parameters
	productIDParam := c.Params("id")
	productID, err := strconv.ParseUint(productIDParam, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product Id "})
	}

	// Retrieve product data from request body
	var updatedProduct model.Product

	if err := c.BodyParser(&updatedProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	// Validate required fields
	if updatedProduct.Name == "" || updatedProduct.Price <= 0 || updatedProduct.ShoppingCategory == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product data",
		})
	}
	// Check if the product exists
	existingProduct, err := utils.FindProductByID((uint(productID)))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	// Update product details
	existingProduct.Name = updatedProduct.Name
	existingProduct.Price = updatedProduct.Price
	existingProduct.ShoppingCategory = updatedProduct.ShoppingCategory

	// Save updated product to database
	if err := database.DB.Save(&existingProduct).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	// Return success response
	return c.JSON(fiber.Map{
		"message": "Product updated successfully",
		"data":    existingProduct,
	})
}

// DeleteProduct deletes product
func DeleteProduct(c *fiber.Ctx) error {

	//Parse product ID from URL parameters
	productIDParam := c.Params("id")
	productID, err := strconv.ParseUint(productIDParam, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product Id "})
	}

	if err := database.DB.Delete(&model.Product{}, productID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	// Return success response
	return c.JSON(fiber.Map{
		"message": "Product delete successfully",
	})
}

// ListProduct list all products
func ListProduct(c *fiber.Ctx) error {

	//Retrive all products from the database
	var products []model.Product

	if err := database.DB.Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch products",
		})
	}

	// Return the list of products as JSON response
	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {

    id, err := c.ParamsInt("id") // Get the product ID from the URL parameter
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid product ID",
        })
    }

    var product model.Product
    result := database.DB.First(&product, id) // Fetch the product from the database

    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Product not found",
        })
    }

    // Return the product data with the updated price
    return c.JSON(product) 
}