package main

import (
	"log"

	"github.com/Sahil-Sayyad/Trademarkia/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	// Create a new Fiber instance
	app := fiber.New()

	//Connect to the database 
	database.ConnectDB()

	// Define a route and its handler
	app.Get("/", func(c *fiber.Ctx) error {
		// Send a string response
		return c.SendString("Hello, World!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
