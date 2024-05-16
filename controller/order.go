package controller

import (
	"github.com/Sahil-Sayyad/Trademarkia/database"
	"github.com/Sahil-Sayyad/Trademarkia/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type OrderInput struct {
    UserID uint   `json:"user_id"`
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

    var user model.User

    if err := database.DB.First(&user, input.UserID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "user not found",
        })
    }

    var order model.Order
    order.UserID = input.UserID

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
    // Load the associated User data using Preload before returning the order
    database.DB.Preload("User").First(&order, order.ID)
    return c.JSON(order)
}


func GetAdminOrders(c *fiber.Ctx) error {

    // Parse query parameters for filtering and sorting
    var filterParams struct {
        UserID    *uint   `query:"user_id"` // Filter by user ID
        ProductID *uint   `query:"product_id"` // Filter by product ID
        Status    *string `query:"status"` // Filter by order status (e.g., "pending", "shipped")
        StartDate *string `query:"start_date"` // Filter by start date (YYYY-MM-DD format)
        EndDate   *string `query:"end_date"` // Filter by end date (YYYY-MM-DD format)
    }

    if err := c.QueryParser(&filterParams); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
    }

    var sortParams struct {
        SortBy  string `query:"sort_by"`  // Field to sort by (e.g., "created_at", "total_price")
        OrderBy string `query:"order_by"` // "asc" or "desc" (default: "desc")
    }
    if err := c.QueryParser(&sortParams); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
    }

    // Build the database query
    var orders []model.Order
    query := database.DB.Preload("Products").Preload("User") // Load related data
    
    if filterParams.UserID != nil {
        query = query.Where("user_id = ?", *filterParams.UserID)
    }
    if filterParams.ProductID != nil {
        query = query.Joins("JOIN order_products ON order_products.order_id = orders.id").
                Where("order_products.product_id = ?", *filterParams.ProductID)
    }
    if filterParams.Status != nil {
        query = query.Where("status = ?", *filterParams.Status)
    }
    if filterParams.StartDate != nil {
        query = query.Where("created_at >= ?", *filterParams.StartDate)
    }
    if filterParams.EndDate != nil {
        query = query.Where("created_at <= ?", *filterParams.EndDate)
    }

    // Apply sorting (use a default if none is provided)
    if sortParams.SortBy == "" {
        sortParams.SortBy = "created_at"
    }
    if sortParams.OrderBy == "" {
        sortParams.OrderBy = "desc"
    }
    query = query.Order(gorm.Expr(sortParams.SortBy + " " + sortParams.OrderBy))
    
    query.Find(&orders) // Execute the query
    return c.JSON(orders)
}

