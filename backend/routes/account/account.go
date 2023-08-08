package account

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router fiber.Router) {
	router.Post("/register/start", registerStart)
	router.Post("/register/verify", registerVerifyEmail)
	router.Post("/register/finish/year", registerYear)
	router.Post("/register/finish/years", registerYearList)
}
