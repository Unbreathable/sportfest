package matches

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type listRequest struct {
	Year uint `json:"year"`
}

// Route: /arq/matches/list
func listMatches(c *fiber.Ctx) error {

	// Parse request
	var req listRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	var games []database.Game
	if err := database.DBConn.Find(&games).Error; err != nil {
		return util.FailedRequest(c, "Es konnten nicht alle Spiele geladen werden.", err)
	}

	// Get matches
	var matches []database.Match
	if req.Year == 0 {
		if err := database.DBConn.Find(&matches).Error; err != nil {
			return util.FailedRequest(c, "Es konnten nicht alle Matches geladen werden.", err)
		}
	} else {
		if err := database.DBConn.Where(&database.Match{Year: req.Year}).Find(&matches).Error; err != nil {
			return util.FailedRequest(c, "Es konnten nicht alle Matches geladen werden.", err)
		}
	}

	return c.JSON(fiber.Map{
		"success": true,
		"matches": matches,
		"games":   games,
	})
}
