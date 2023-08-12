package algorithm

import (
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

func accept(c *fiber.Ctx) error {

	// Set teams algorithm chose

	return util.SuccessfulRequest(c)
}
