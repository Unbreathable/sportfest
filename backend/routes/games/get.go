package years

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type getRequest struct {
	ID uint `json:"id"`
}

// Route: /arq/games/get
func getGame(c *fiber.Ctx) error {

	// Parse request
	var req getRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	if req.ID == 0 {
		return util.InvalidRequest(c)
	}

	var game database.Game
	if database.DBConn.Where(&database.Game{ID: req.ID}).Take(&game).Error != nil {
		return util.FailedRequest(c, "Dieses Jahr existiert nicht!", nil)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"obj":     game,
	})
}
