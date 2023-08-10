package years

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type deleteRequest struct {
	ID uint `json:"id"`
}

// Route: /arq/years/delete
func deleteYear(c *fiber.Ctx) error {

	// Parse request
	var req deleteRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	if req.ID == 0 {
		return util.InvalidRequest(c)
	}

	// Delete everything related to this year
	if database.DBConn.Where(&database.User{Year: req.ID}).Delete(&database.User{}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Löschen aller Benutzer in diesem Jahr.", nil)
	}

	if database.DBConn.Where(&database.Match{Year: req.ID}).Delete(&database.Match{}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Löschen aller Spielewahlen in diesem Jahr.", nil)
	}

	if database.DBConn.Where(&database.Year{ID: req.ID}).Delete(&database.Year{}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Löschen dieses Jahres.", nil)
	}

	return util.SuccessfulRequest(c)
}
