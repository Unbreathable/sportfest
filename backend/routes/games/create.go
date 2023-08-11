package games

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type createGameRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Video       string `json:"video"`
	MinTeamSize uint   `json:"minTeamSize"`
	MaxTeamSize uint   `json:"maxTeamSize"`
}

// Route: /arq/games/create
func createGame(c *fiber.Ctx) error {

	// Parse request
	var req createGameRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	if req.MinTeamSize == 0 || req.MaxTeamSize == 0 {
		return util.FailedRequest(c, "Bitte gib eine Mindest- und eine Maximalgröße an!", nil)
	}

	if req.MinTeamSize > req.MaxTeamSize {
		return util.FailedRequest(c, "Die Mindestgröße darf nicht größer als die Maximalgröße sein!", nil)
	}

	if req.Name == "" {
		return util.FailedRequest(c, "Bitte gib einen Namen an!", nil)
	}

	if len(req.Description) < 10 {
		return util.FailedRequest(c, "Die Beschreibung muss mindestens 10 Zeichen lang sein!", nil)
	}

	if len(req.Name) < 4 {
		return util.FailedRequest(c, "Der Name muss mindestens 4 Zeichen lang sein!", nil)
	}

	var game database.Game
	if database.DBConn.Where("name = ?", req.Name).Take(&game).Error == nil {
		return util.FailedRequest(c, "Dieses Jahr existiert bereits!", nil)
	}

	// Create year
	game = database.Game{
		Name:        req.Name,
		Description: req.Description,
		MinTeamSize: req.MinTeamSize,
		MaxTeamSize: req.MaxTeamSize,
	}
	database.DBConn.Create(&game)

	return c.JSON(fiber.Map{
		"success": true,
		"obj":     game.ID,
	})
}
