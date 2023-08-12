package friends

import "github.com/gofiber/fiber/v2"

func Unauthorized(router fiber.Router) {
	router.Post("/status", status)
	router.Post("/generate", generate)
	router.Post("/set", set)
}

func Authorized(router fiber.Router) {
	router.Post("/simulation", friendsSimulation)
}
