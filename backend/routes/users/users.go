package users

import (
	"github.com/Unbreathable/sportfest/backend/routes/users/friends"
	"github.com/gofiber/fiber/v2"
)

const MinChoices = 3
const MaxChoices = 5

func Unauthorized(router fiber.Router) {
	router.Post("/choices", choices)
	router.Post("/removechoice", removeChoice)
	router.Post("/addchoice", addChoice)
	router.Route("/friends", friends.Unauthorized)
}

func Authorized(router fiber.Router) {
	router.Post("/simulation", simulation)
	router.Post("/choicesimulation", choiceSimulation)
	router.Post("/search", search)
	router.Route("/friends", friends.Authorized)
}
