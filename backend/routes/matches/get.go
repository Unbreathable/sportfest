package matches

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type getRequest struct {
	Match uint `json:"match"`
}

// Route: /arq/matches/get
func getMatch(c *fiber.Ctx) error {

	// Parse request
	var req getRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	var match database.Match
	if err := database.DBConn.Where(&database.Match{ID: req.Match}).First(&match).Error; err != nil {
		return util.FailedRequest(c, "Das Match konnte nicht geladen werden.", err)
	}

	var game database.Game
	if err := database.DBConn.Where(&database.Game{ID: match.Game}).First(&game).Error; err != nil {
		return util.FailedRequest(c, "Das Spiel konnte nicht geladen werden.", err)
	}

	var team1 []database.User
	if database.DBConn.Raw("SELECT * FROM users WHERE EXISTS (SELECT * FROM team_members WHERE team_members.user = users.id AND team_members.match = ? AND team_members.team = ?)", match.ID, database.Team1).Scan(&team1).Error != nil {
		return util.FailedRequest(c, "Team 1 konnte nicht geladen werden.", nil)
	}

	var team2 []database.User
	if database.DBConn.Raw("SELECT * FROM users WHERE EXISTS (SELECT * FROM team_members WHERE team_members.user = users.id AND team_members.match = ? AND team_members.team = ?)", match.ID, database.Team2).Scan(&team2).Error != nil {
		return util.FailedRequest(c, "Team 2 konnte nicht geladen werden.", nil)
	}

	// TODO: Compute suggestions

	return c.JSON(fiber.Map{
		"success": true,
		"match":   match,
		"game":    game,
		"team1":   team1,
		"team2":   team2,
	})
}
