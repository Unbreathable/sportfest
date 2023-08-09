package account

import "github.com/gofiber/fiber/v2"

// Unauthorized routes
func SetupRoutes(router fiber.Router) {

	router.Post("/login", login)

	// We'll probably not end up using this
	router.Post("/register/start", registerStart)
	router.Post("/register/verify", registerVerifyEmail)
	router.Post("/register/finish/year", registerYear)
	router.Post("/register/finish/years", registerYearList)
}

// Authorized routes
func Authorized(router fiber.Router) {
	router.Post("/create", createAdminAccount)
}
