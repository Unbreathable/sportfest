package users

import (
	"fmt"

	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type simulationRequest struct {
	MinUsersYear int `json:"minPerYear"`
	MaxUsersYear int `json:"maxPerYear"`
}

// Route: /arq/users/simulation
func simulation(c *fiber.Ctx) error {

	// Parse request
	var req simulationRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	var years []database.Year
	if database.DBConn.Find(&years).Error != nil {
		return util.InvalidRequest(c)
	}

	if req.MinUsersYear < 0 || req.MaxUsersYear < 0 || req.MinUsersYear > req.MaxUsersYear {
		return util.InvalidRequest(c)
	}

	// Create users
	for _, year := range years {

		amount := util.RandomInt(req.MinUsersYear, req.MaxUsersYear)
		for i := 0; i < amount; i++ {
			user := database.User{
				Name: fmt.Sprintf("%d-%s", year.ID, util.GenerateToken(6)),
				Code: util.GenerateToken(12),
				Year: year.ID,
			}
			if err := database.DBConn.Create(&user).Error; err != nil {
				return util.InvalidRequest(c)
			}
		}
	}

	return util.SuccessfulRequest(c)
}
