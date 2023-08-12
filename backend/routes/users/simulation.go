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

func choiceSimulation(c *fiber.Ctx) error {

	var users []database.User
	if database.DBConn.Find(&users).Error != nil {
		return util.InvalidRequest(c)
	}

	var games []database.Game
	if database.DBConn.Find(&games).Error != nil {
		return util.InvalidRequest(c)
	}

	for _, user := range users {

		// Generate random amount of choices
		amount := util.RandomInt(MinChoices, MaxChoices+1)
		var chosen []uint
		for i := 0; i < amount; i++ {

			// Generate random choice
			game := games[util.RandomInt(0, len(games))].ID
			if contains(chosen, game) {
				i--
				continue
			}
			chosen = append(chosen, game)

			// Create choice
			choice := database.Choice{
				User:   user.ID,
				Choice: game,
			}
			if err := database.DBConn.Create(&choice).Error; err != nil {
				return util.InvalidRequest(c)
			}
		}
	}

	return util.SuccessfulRequest(c)
}

func contains(arr []uint, val uint) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
