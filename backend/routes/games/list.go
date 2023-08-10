package years

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

// Route: /arq/games/list
func listGames(c *fiber.Ctx) error {

	var games []database.Game
	if database.DBConn.Find(&games).Error != nil {
		return util.FailedRequest(c, "Fehler beim Laden der auswählbaren Spiele!", nil)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"obj":     games,
	})
}
