package games

import "github.com/gofiber/fiber/v2"

func Authorized(router fiber.Router) {
	router.Post("/create", createGame)
	router.Post("/list", listGames)
	router.Post("/get", getGame)
	router.Post("/delete", deleteGame)
	router.Post("/default", createDefault)
}
