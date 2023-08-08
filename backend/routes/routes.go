package routes

import (
	"log"

	"github.com/Unbreathable/sportfest/backend/util"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {

	// Unauthorized routes

	// Setup authorized routes
	router.Route("/", AuthorizedRoutes)
}

func AuthorizedRoutes(router fiber.Router) {

	// Authorized using JWT
	router.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(util.JWT_SECRET),
		},

		// Checks if the token is expired
		SuccessHandler: func(c *fiber.Ctx) error {

			if util.IsExpired(c) {
				return util.InvalidRequest(c)
			}

			return c.Next()
		},

		// Error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			log.Println(err.Error())

			// Return error message
			return c.SendStatus(401)
		},
	}))

}
