package main

import (
	"log"
	"os"

	"github.com/Unbreathable/sportfest/backend/caching"
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/routes"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	util.JWT_SECRET = os.Getenv("JWT_SECRET")

	// Setup data storage
	database.Setup()
	caching.SetupCaches()

	// Start the fiber server
	app := fiber.New()
	app.Route("/", routes.SetupRoutes)

	// Start the server on port 3000 (in prod bind to 0.0.0.0)
	app.Listen("localhost:3000")
}
