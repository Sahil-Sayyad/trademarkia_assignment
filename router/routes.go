package router

import (
	"github.com/Sahil-Sayyad/Trademarkia/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){

	api := app.Group("/api")

	//User Sign-up
	api.Post("/users", controller.CreateUser)

}