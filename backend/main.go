package main

import (
	"log"
	"os"

	"github.com/Unbreathable/sportfest/backend/caching"
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/routes"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	createDefaultAccount()

	// Start the fiber server
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	app.Route("/", routes.SetupRoutes)

	// Start the server on port 3000 (in prod bind to 0.0.0.0)
	app.Listen("localhost:3000")
}

func createDefaultAccount() {

	var acc database.AdminAccount
	if database.DBConn.Where("name = ?", "admin").Take(&acc).Error != nil {
		acc.Name = "admin"
		acc.Password = util.GenerateToken(12)
		log.Println("Created default admin account with password:", acc.Password)

		acc.Password = util.HashPassword(acc.Password)
		database.DBConn.Create(&acc)
	}

}
