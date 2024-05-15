package router

import (
	"github.com/Sahil-Sayyad/Trademarkia/controller"
	"github.com/Sahil-Sayyad/Trademarkia/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){

	api := app.Group("/api")

	//User Sign-up
	api.Post("/users/sign-up", controller.CreateUser)
	//User login 
	api.Post("/users/login", controller.Login)

	admin := api.Group("/admin")

	//Admin Sign-up
	admin.Post("/sign-up", controller.CreateAdmin)
	//Admin Login
	admin.Post("/login", controller.LoginAdmin)

    
	admin.Use(middleware.AdminauthMiddleware)
    //Admin Adding Product (admin access only)
	admin.Post("/products", controller.CreateProduct)
	admin.Put("/products/:id", controller.UpdateProduct)
}