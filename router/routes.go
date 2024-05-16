package router

import (
	"github.com/Sahil-Sayyad/Trademarkia/controller"
	"github.com/Sahil-Sayyad/Trademarkia/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// User-Side APIs:
	api := app.Group("/api")
	// POST /api/users/signup  (Create a new user)
	api.Post("/users/sign-up", controller.CreateUser)
	// POST /api/users/login   (User authentication)
	api.Post("/users/login", controller.Login)
	// GET  /api/products      (Search for products)
	api.Get("/products", controller.ListProduct)
	// GET  /api/products/:id  (Get a specific product)
	api.Get("/products/:id", controller.GetProduct)
	// POST /api/orders        (Place an order)
	api.Post("/orders",middleware.UserauthMiddleware, controller.CreateOrder)
	// GET  /api/users/dashboard (View user's order history)
	api.Get("/users/dashboard", middleware.UserauthMiddleware, controller.GetUserDashboard)

	// Admin-Side APIs:
	admin := api.Group("/admin")
	// POST   /api/admin/sign-up      (create a new admin)
	admin.Post("/sign-up", controller.CreateAdmin)
	// POST   /api/admin/login        (login admin)
	admin.Post("/login", controller.LoginAdmin)
	// Admin auth middleware for admin access only
	admin.Use(middleware.AdminauthMiddleware)
	// POST   /api/admin/products     (Add a new product)
	admin.Post("/products", controller.CreateProduct)
	// PUT    /api/admin/products/:id (Update a product)
	admin.Put("/products/:id", controller.UpdateProduct)
	// DELETE /api/admin/products/:id (Remove a product)
	admin.Delete("/products/:id", controller.DeleteProduct)
	// GET    /api/admin/orders       (Get all orders with filters/sorting)
	admin.Get("/orders", controller.GetAdminOrders)
	// GET    /api/admin/stats        (Get statistics on orders, inventory)
	admin.Get("/stats", controller.GetAdminStats)
}
