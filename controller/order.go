
package controller

import (
	"github.com/Sahil-Sayyad/Trademarkia/database"
	"github.com/Sahil-Sayyad/Trademarkia/model"
	"github.com/gofiber/fiber/v2"
)

type OrderInput struct {
    CustomerID uint   `json:"customer_id"`
    Products   []struct {
        ProductID uint `json:"product_id"`
        Quantity  uint `json:"quantity"`
    } `json:"products"`
}

func CreateOrder(c *fiber.Ctx) error {
    var input OrderInput
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input data",
        })
    }

    var customer model.User
    if err := database.DB.First(&customer, input.CustomerID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Customer not found",
        })
    }

    var order model.Order
    order.CustomerID = input.CustomerID

    var totalPrice float64
    for _, item := range input.Products {
        var product model.Product
        if err := database.DB.First(&product, item.ProductID).Error; err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "Product not found",
            })
        }

        if product.Quantity < int(item.Quantity) {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Insufficient product quantity",
            })
        }

        product.Quantity -= int(item.Quantity)
        database.DB.Save(&product)

        order.Products = append(order.Products, product)
        totalPrice += product.Price * float64(item.Quantity)
    }

    order.TotalPrice = totalPrice
    database.DB.Create(&order)

    return c.JSON(order)
}
