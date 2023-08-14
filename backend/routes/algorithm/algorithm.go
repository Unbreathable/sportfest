package algorithm

import "github.com/gofiber/fiber/v2"

func Authorized(router fiber.Router) {
	router.Post("/start", start)
	router.Post("/accept", accept)
}
