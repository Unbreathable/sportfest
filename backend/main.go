package main

import (
	"log"

	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	database.Setup()

	// Start a basic fiber server
	app := fiber.New()

	// Create a GET route on path "/hello"
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start the server on port 3000
	app.Listen("localhost:3000")
}
