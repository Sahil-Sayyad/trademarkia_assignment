package router

import (
	"github.com/gofiber/fiber/v2"

)

func SetupRoutes(app *fiber.App){
	api := app.Group("/api")

	api.Get("/users", func(c *fiber.Ctx) error {
		return c.SendString("Done")
	})
}