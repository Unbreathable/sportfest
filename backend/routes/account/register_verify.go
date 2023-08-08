package account

import (
	"github.com/Unbreathable/sportfest/backend/caching"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

const MaxAttempts = 10

type registerVerifyRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func registerVerifyEmail(c *fiber.Ctx) error {

	// Parse request
	var req registerVerifyRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Normalize email
	email := NormalizeEmail(req.Email)

	// Check if code is valid
	token, valid := caching.VerifyEmail(email, req.Code)
	if !valid {
		return util.FailedRequest(c, "Der Code ist ungültig!", nil)
	}

	// Check if user has tried too often
	if token.Attempts > MaxAttempts {
		caching.RemoveEmail(email)
		return util.FailedRequest(c, "Du hast zu oft versucht den Code einzugeben!", nil)
	}

	return c.JSON(fiber.Map{
		"success":      true,
		"attempts":     token.Attempts,
		"max_attempts": MaxAttempts,
	})
}
