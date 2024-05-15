package controller

import (
	"github.com/Sahil-Sayyad/Trademarkia/database"
	"github.com/Sahil-Sayyad/Trademarkia/model"
	"github.com/gofiber/fiber/v2"
)

func createProduct(c *fiber.Ctx) error {

	var product model.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	//Create user in the database
	result := database.DB.Create(&product)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create product"})
	}

	return c.JSON(product)
}
