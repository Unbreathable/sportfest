package years

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

// Route: /arq/games/default
func createDefault(c *fiber.Ctx) error {

	// Create default games
	if database.DBConn.Create(&database.Game{
		Name:        "Völkerball",
		Description: "Ziel des Spiels ist es, alle Spieler der gegnerischen Mannschaft durch Treffen mit einem Ball aus dem Spiel zu werfen oder zu fangen, ohne selbst getroffen zu werden oder den Ball zu fangen.",
		MinTeamSize: 5,
		MaxTeamSize: 10,
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardspiels.", nil)
	}
	if database.DBConn.Create(&database.Game{
		Name:        "Brennball",
		Description: "Ziel des Spiels ist es, den Ball öfter als die gegnerische Mannschaft in das gegnerische Tor zu befördern.",
		MinTeamSize: 7,
		MaxTeamSize: 20,
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardspiels.", nil)
	}
	if database.DBConn.Create(&database.Game{
		Name:        "Battle Box",
		Description: "Ziel des Spiels ist es, den Ball öfter als die gegnerische Mannschaft in das gegnerische Tor zu befördern.",
		MinTeamSize: 4,
		MaxTeamSize: 7,
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardspiels.", nil)
	}
	if database.DBConn.Create(&database.Game{
		Name:        "Staffellauf",
		Description: "Ziel des Spiels ist es, den Ball öfter als die gegnerische Mannschaft in das gegnerische Tor zu befördern.",
		MinTeamSize: 3,
		MaxTeamSize: 6,
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardspiels.", nil)
	}
	if database.DBConn.Create(&database.Game{
		Name:        "Hindernislauf",
		Description: "Ziel des Spiels ist es, den Ball öfter als die gegnerische Mannschaft in das gegnerische Tor zu befördern.",
		MinTeamSize: 3,
		MaxTeamSize: 6,
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardspiels.", nil)
	}
	if database.DBConn.Create(&database.Game{
		Name:        "Reiterkampf",
		Description: "Ziel des Spiels ist es, den Ball öfter als die gegnerische Mannschaft in das gegnerische Tor zu befördern.",
		MinTeamSize: 6,
		MaxTeamSize: 12,
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardspiels.", nil)
	}
	if database.DBConn.Create(&database.Game{
		Name:        "Capture The Cone",
		Description: "Ziel des Spiels ist es, den Ball öfter als die gegnerische Mannschaft in das gegnerische Tor zu befördern.",
		MinTeamSize: 7,
		MaxTeamSize: 20,
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardspiels.", nil)
	}

	return util.SuccessfulRequest(c)
}
