package account

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type simulationRequest struct {
	MinUsersYear int `json:"min_users_year"`
	MaxUsersYear int `json:"max_users_year"`
	Friendships  int `json:"friendships"`
}

// Route: /arq/account/simulation
func simulation(c *fiber.Ctx) error {

	// Parse request
	var req simulationRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	var years []uint
	if database.DBConn.Model(&database.Year{}).Select("id").Find(&years).Error != nil {
		return util.InvalidRequest(c)
	}

	// TODO: Finish simulation

	return util.SuccessfulRequest(c)
}
