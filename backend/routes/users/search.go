package users

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type searchRequest struct {
	Query string `json:"query"`
}

// Route: /arq/users/search
func search(c *fiber.Ctx) error {

	// Parse request
	var req searchRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	var users []database.User

	if req.Query == "" {
		if database.DBConn.Limit(20).Find(&users).Error != nil {
			return util.InvalidRequest(c)
		}
		return c.JSON(fiber.Map{
			"success": true,
			"obj":     users,
		})
	} else if database.DBConn.Where("name LIKE ?", "%"+req.Query+"%").Limit(20).Find(&users).Error != nil {
		return c.JSON(fiber.Map{
			"success": true,
			"obj":     []interface{}{},
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"obj":     users,
	})
}
