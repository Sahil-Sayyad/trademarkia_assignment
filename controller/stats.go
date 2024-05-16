package controller

import (
	"github.com/Sahil-Sayyad/Trademarkia/database"
	"github.com/Sahil-Sayyad/Trademarkia/model"
	"github.com/gofiber/fiber/v2"
)



type StatsResponse struct {
    TotalOrders          int64         `json:"total_orders"`
    TotalRevenue         float64      `json:"total_revenue"`
    OrdersByStatus       map[string]uint `json:"orders_by_status"` // e.g., { "pending": 10, "shipped": 25, ... }
    LowStockProducts     []model.Product `json:"low_stock_products"`
    TotalUsers           int64         `json:"total_users"`
    RecentOrders         []model.Order   `json:"recent_orders"`  // Last 10 orders (for example)
    AverageOrderValue    float64      `json:"average_order_value"`
}

func GetAdminStats(c *fiber.Ctx) error {
  
    var stats StatsResponse

    // Total Orders
    database.DB.Model(&model.Order{}).Count(&stats.TotalOrders)

    // Total Revenue
    database.DB.Model(&model.Order{}).Select("SUM(total_price)").Scan(&stats.TotalRevenue)

    // Orders by Status
    stats.OrdersByStatus = make(map[string]uint)
    database.DB.Model(&model.Order{}).Select("status, COUNT(*)").Group("status").Scan(&stats.OrdersByStatus)

    // Low Stock Products (e.g., quantity < 10)
    database.DB.Where("quantity < ?", 10).Find(&stats.LowStockProducts)

    // Total Users
    database.DB.Model(&model.User{}).Count(&stats.TotalUsers)

    // Recent Orders (last 10)
    database.DB.Preload("Products").Order("created_at desc").Limit(10).Find(&stats.RecentOrders)

    // Average Order Value
    if stats.TotalOrders > 0 {
        stats.AverageOrderValue = stats.TotalRevenue / float64(stats.TotalOrders)
    }

    return c.JSON(stats)
}
