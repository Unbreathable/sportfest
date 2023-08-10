package years

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

// Route: /arq/years/default
func createDefault(c *fiber.Ctx) error {

	// Create default years
	if database.DBConn.Create(&database.Year{
		Name: "5. Klasse",
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardjahres.", nil)
	}
	if database.DBConn.Create(&database.Year{
		Name: "6. Klasse",
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardjahres.", nil)
	}
	if database.DBConn.Create(&database.Year{
		Name: "7. Klasse",
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardjahres.", nil)
	}
	if database.DBConn.Create(&database.Year{
		Name: "8. Klasse",
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardjahres.", nil)
	}
	if database.DBConn.Create(&database.Year{
		Name: "9. Klasse",
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardjahres.", nil)
	}
	if database.DBConn.Create(&database.Year{
		Name: "10. Klasse",
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardjahres.", nil)
	}
	if database.DBConn.Create(&database.Year{
		Name: "Kursstufe 1",
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardjahres.", nil)
	}
	if database.DBConn.Create(&database.Year{
		Name: "Kursstufe 2",
	}).Error != nil {
		return util.FailedRequest(c, "Es gab einen Fehler beim Erstellen des Standardjahres.", nil)
	}

	return util.SuccessfulRequest(c)
}
