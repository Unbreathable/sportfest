package friends

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type generateRequest struct {
	Code string `json:"code"`
}

// Route: /users/friends/generate
func generate(c *fiber.Ctx) error {

	// Parse request
	var req generateRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Get user
	var user database.User
	if database.DBConn.Where(&database.User{Code: req.Code}).Take(&user).Error != nil {
		return util.InvalidRequest(c)
	}

	// Check if user already has a friend code
	var friendship database.Friendship
	if database.DBConn.Where(&database.Friendship{ID: user.ID}).Take(&friendship).Error == nil {
		return util.FailedRequest(c, "Du hast bereits einen Freundescode generiert.", nil)
	}

	// Generate friend code
	friendship = database.Friendship{
		ID:   user.ID,
		Code: util.GenerateToken(16),
	}
	if database.DBConn.Save(&friendship).Error != nil {
		return util.FailedRequest(c, "Ein unerwarteter Fehler ist aufgetreten, bitte versuche es später erneut.", nil)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"code":    friendship.Code,
	})
}
