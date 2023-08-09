package account

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type loginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func login(c *fiber.Ctx) error {

	// Parse request
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Get account
	var acc database.AdminAccount
	if database.DBConn.Where("name = ?", req.Name).Take(&acc).Error != nil {
		return util.FailedRequest(c, "Dein Benutzername oder Passwort ist falsch.", nil)
	}

	// Check password
	if util.HashPassword(req.Password) != acc.Password {
		return util.FailedRequest(c, "Dein Benutzername oder Passwort ist falsch.", nil)
	}

	// Generate token
	token, err := util.Token(acc.ID)
	if err != nil {
		return util.FailedRequest(c, "Etwas lief auf dem Server schief. Bitte benachrichtige einen Administrator über diesen Fehler.", nil)
	}

	// Return token
	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}
