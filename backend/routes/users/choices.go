package users

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type choicesRequest struct {
	Code string `json:"code"`
}

func choices(c *fiber.Ctx) error {

	// Parse request
	var req choicesRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Get user
	var user database.User
	if database.DBConn.Where(&database.User{Code: req.Code}).Take(&user).Error != nil {
		return util.FailedRequest(c, "Du hast keinen Zugriff auf diese Seite!", nil)
	}

	// Get games and choices
	var games []database.Game
	var choices []uint
	if err := database.DBConn.Find(&games).Error; err != nil {
		return util.FailedRequest(c, "Fehler beim Laden der Spiele!", nil)
	}
	if err := database.DBConn.Model(&database.Choice{}).Select("choice").Where(&database.Choice{User: user.ID}).Find(&choices).Error; err != nil {
		return util.FailedRequest(c, "Fehler beim Laden der Auswahl!", nil)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"games":   games,
		"choices": choices,
	})
}
