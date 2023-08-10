package account

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type createAdminAccountRequest struct {
	Name string `json:"name"`
}

// Route: /arq/account/create
func createAdminAccount(c *fiber.Ctx) error {

	// Parse request
	var req createAdminAccountRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Generate account
	password := util.GenerateToken(12)
	acc := database.AdminAccount{
		Name:     req.Name,
		Password: util.HashPassword(password),
	}
	if database.DBConn.Create(&acc).Error != nil {
		return util.InvalidRequest(c)
	}

	token, err := util.Token(acc.ID)
	if err != nil {
		return util.InvalidRequest(c)
	}

	// Return password and token
	return c.JSON(fiber.Map{
		"success":  true,
		"password": password,
		"token":    token,
	})
}
