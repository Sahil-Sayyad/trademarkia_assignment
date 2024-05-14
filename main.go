package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Define a route and its handler
	app.Get("/", func(c *fiber.Ctx) error {
		// Send a string response
		return c.SendString("Hello, World!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
