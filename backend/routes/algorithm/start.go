package algorithm

import (
	"fmt"
	"log"

	"github.com/Unbreathable/sportfest/backend/algorithm"
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

var running = false
var algGames = []algorithm.Game{}
var algPlayers = []*algorithm.Player{}
var algYears = []algorithm.Year{}

func start(c *fiber.Ctx) error {
	if running {
		return util.FailedRequest(c, "Der Algorithmus läuft bereits.", nil)
	}
	running = true
	defer func() {
		running = false
	}()

	// Get years, games and players
	var years []database.Year
	var games []database.Game
	var users []database.User

	if database.DBConn.Find(&years).Error != nil {
		return util.FailedRequest(c, "Es konnten nicht alle Jahrgänge geladen werden.", nil)
	}
	if database.DBConn.Find(&games).Error != nil {
		return util.FailedRequest(c, "Es konnten nicht alle Spiele geladen werden.", nil)
	}
	if database.DBConn.Find(&users).Error != nil {
		return util.FailedRequest(c, "Es konnten nicht alle Benutzer geladen werden.", nil)
	}

	// Convert users to algorithm types
	algPlayers = []*algorithm.Player{}
	for _, user := range users {
		algPlayer := algorithm.Player{
			ID:   user.ID,
			Year: user.Year,
		}

		var choices []uint
		if err := database.DBConn.Model(&database.Choice{}).Select("choice").Where(&database.Choice{User: user.ID}).Find(&choices).Error; err != nil {
			return util.FailedRequest(c, fmt.Sprintf("Die Auswahl des Benutzers %s (ID: %d) konnte nicht geladen werden.", user.Name, user.ID), err)
		}
		algPlayer.GamesSelected = choices
		algPlayer.Friend = GetFriendship(user.ID)

		algPlayers = append(algPlayers, &algPlayer)
	}

	// Convert years to algorithm types
	algYears = []algorithm.Year{}
	for _, year := range years {

		usersInYear := []*algorithm.Player{}
		for _, user := range algPlayers {
			if user.Year == year.ID {
				usersInYear = append(usersInYear, user)
			}
		}

		algYear := algorithm.Year{
			ID:         year.ID,
			Player:     usersInYear,
			GameAmount: map[uint]int{},
			Playable:   map[uint]*algorithm.PlayableGame{},
		}
		algYears = append(algYears, algYear)
	}

	// Convert games to algorithm types
	algGames = []algorithm.Game{}
	for _, game := range games {
		algGames = append(algGames, algorithm.Game{
			ID:          game.ID,
			MinTeamSize: int(game.MinTeamSize),
			MaxTeamSize: int(game.MaxTeamSize),
		})
	}

	logs, valid := algorithm.Compute(&algYears, &algGames, &algPlayers)

	if valid {
		return c.JSON(fiber.Map{
			"success":   true,
			"algorithm": false,
			"message":   "",
			"logs":      logs,
		})
	}

	return c.JSON(fiber.Map{
		"success":   true,
		"algorithm": false,
		"message":   "Der Algorithmus konnte nicht erfolgreich ausgeführt werden.",
		"logs":      logs,
	})
}

func GetFriendship(user uint) uint {

	var friendship database.Friendship
	if database.DBConn.Where(&database.Friendship{ID: user}).Take(&friendship).Error != nil {
		return 0
	}

	// Check if friendship is valid
	var friendshipCheck database.Friendship
	if database.DBConn.Where(&database.Friendship{Code: friendship.FriendCode}).Take(&friendshipCheck).Error != nil {
		return 0
	}

	if friendship.Code != friendshipCheck.FriendCode {
		return 0
	}

	log.Println("Vaild friendship found!")

	return friendshipCheck.ID
}
