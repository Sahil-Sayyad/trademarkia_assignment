package router

import (
	"github.com/Sahil-Sayyad/Trademarkia/controller"
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
    //Admin Adding Product 
	admin.Post("/products", controller.createProduct)
}