package users

import "github.com/gofiber/fiber/v2"

const MinChoices = 3
const MaxChoices = 5

func Unauthorized(router fiber.Router) {
	router.Post("/choices", choices)
	router.Post("/removechoice", removeChoice)
	router.Post("/addchoice", addChoice)
}

func Authorized(router fiber.Router) {
	router.Post("/simulation", simulation)
	router.Post("/search", search)
}
