package games

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type deleteRequest struct {
	ID uint `json:"id"`
}

// Route: /arq/games/delete
func deleteGame(c *fiber.Ctx) error {

	// Parse request
	var req deleteRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	if req.ID == 0 {
		return util.InvalidRequest(c)
	}

	// Delete everything related to this game
	if database.DBConn.Where(&database.Choice{Choice: req.ID}).Delete(&database.Choice{}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Löschen aller Auswahlen dieses Spiels.", nil)
	}

	if database.DBConn.Where(&database.Match{Game: req.ID}).Delete(&database.Match{}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Löschen aller Matches dieses Spiels.", nil)
	}

	if database.DBConn.Where(&database.TeamMember{Game: req.ID}).Delete(&database.TeamMember{}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Löschen aller Teammitglieder dieses Spiels.", nil)
	}

	if database.DBConn.Where(&database.Game{ID: req.ID}).Delete(&database.Game{}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Löschen dieses Spiels.", nil)
	}

	return util.SuccessfulRequest(c)
}
