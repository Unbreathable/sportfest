package algorithm

import (
	"github.com/Unbreathable/sportfest/backend/algorithm"
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

// Route: /arq/algorithm/accept
func accept(c *fiber.Ctx) error {

	if len(algYears) == 0 {
		return util.FailedRequest(c, "Es wurde noch kein Algorithmus gestartet.", nil)
	}

	// Set teams algorithm chose
	for _, year := range algYears {

		// Create matches
		for _, game := range year.Playable {
			for _, match := range game.AvailableMatches {
				dbMatch := database.Match{
					Game:     game.ID,
					Year:     year.ID,
					TeamSize: match.TeamSize,
				}
				database.DBConn.Create(&dbMatch)

				// Set teams for match
				for _, team := range match.Team1 {
					dbMember := database.TeamMember{
						Team:  database.Team1,
						Match: dbMatch.ID,
						Game:  game.ID,
						User:  team.ID,
					}
					database.DBConn.Create(&dbMember)
					database.DBConn.Model(&database.User{}).Where("id = ?", team.ID).Update("team", database.Team1)
				}

				for _, team := range match.Team2 {
					dbMember := database.TeamMember{
						Team:  database.Team2,
						Match: dbMatch.ID,
						Game:  game.ID,
						User:  team.ID,
					}
					database.DBConn.Create(&dbMember)
					database.DBConn.Model(&database.User{}).Where("id = ?", team.ID).Update("team", database.Team2)
				}

			}
		}

	}

	// Reset algorithm
	algYears = []algorithm.Year{}
	algGames = []algorithm.Game{}
	algPlayers = []*algorithm.Player{}

	return util.SuccessfulRequest(c)
}
