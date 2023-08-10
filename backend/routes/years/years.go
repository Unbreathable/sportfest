package years

import "github.com/gofiber/fiber/v2"

func Authorized(router fiber.Router) {
	router.Post("/create", createYear)
	router.Post("/list", listYears)
}
