package routes

import (
	"log"

	"github.com/Unbreathable/sportfest/backend/routes/account"
	"github.com/Unbreathable/sportfest/backend/routes/games"
	"github.com/Unbreathable/sportfest/backend/routes/users"
	"github.com/Unbreathable/sportfest/backend/routes/years"
	"github.com/Unbreathable/sportfest/backend/util"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {

	// Unauthorized routes
	router.Route("/account", account.SetupRoutes)
	router.Route("/users", users.Unauthorized)

	// Setup authorized routes
	router.Route("/arq", AuthorizedRoutes)
}

func AuthorizedRoutes(router fiber.Router) {

	// Authorized using JWT
	router.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(util.JWT_SECRET),
		},

		// We don't do anything here since we don't care that much about security
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},

		// Error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			log.Println(c.Route().Path + " | " + err.Error())

			// Return error message
			return c.SendStatus(401)
		},
	}))

	// Setup routes
	router.Route("/account", account.Authorized)
	router.Route("/years", years.Authorized)
	router.Route("/games", games.Authorized)
	router.Route("/users", users.Authorized)
}
