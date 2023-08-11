package users

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type chooseRequest struct {
	Code   string `json:"code"`
	Choice uint   `json:"choice"`
}

// Route: /users/addchoice
func addChoice(c *fiber.Ctx) error {

	// Parse request
	var req chooseRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Get user
	var user database.User
	if database.DBConn.Where(&database.User{Code: req.Code}).Take(&user).Error != nil {
		return util.FailedRequest(c, "Du hast keinen Zugriff auf diese Seite!", nil)
	}

	// Get count of choices
	var count int64
	if err := database.DBConn.Model(&database.Choice{}).Where(&database.Choice{User: user.ID}).Count(&count).Error; err != nil {
		return util.FailedRequest(c, "Fehler beim Laden der Auswahl!", nil)
	}

	if count >= MaxChoices {
		return util.FailedRequest(c, "Du hast bereits 5 Spiele ausgewählt. Lösche zuerst die Auswahl von einem anderen Spiel, bevor du wieder ein neues auswählen kannst.", nil)
	}

	// Check if choice is valid
	var game database.Game
	if database.DBConn.Where(&database.Game{ID: req.Choice}).Take(&game).Error != nil {
		return util.FailedRequest(c, "Dieses Spiel existiert nicht!", nil)
	}

	// Check if choice is already made
	var choice database.Choice
	if database.DBConn.Where(&database.Choice{User: user.ID, Choice: req.Choice}).Take(&choice).Error == nil {
		return util.FailedRequest(c, "Du hast dieses Spiel bereits ausgewählt!", nil)
	}

	// Add choice
	if err := database.DBConn.Create(&database.Choice{User: user.ID, Choice: req.Choice}).Error; err != nil {
		return util.FailedRequest(c, "Fehler beim Hinzufügen der Auswahl!", nil)
	}

	return util.SuccessfulRequest(c)
}

// Route: /users/removechoice
func removeChoice(c *fiber.Ctx) error {

	// Parse request
	var req chooseRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Get user
	var user database.User
	if database.DBConn.Where(&database.User{Code: req.Code}).Take(&user).Error != nil {
		return util.FailedRequest(c, "Du hast keinen Zugriff auf diese Seite!", nil)
	}

	// Get count of choices
	var count int64
	if err := database.DBConn.Model(&database.Choice{}).Where(&database.Choice{User: user.ID}).Count(&count).Error; err != nil {
		return util.FailedRequest(c, "Fehler beim Laden der Auswahl!", nil)
	}

	if count <= MinChoices {
		return util.FailedRequest(c, "Du musst mindestens 3 Spiele auswählen. Wähle zuerst ein anderes Spiel aus, bevor eine Auswahl wieder löschen kannst.", nil)
	}

	// Check if choice is already made
	var choice database.Choice
	if database.DBConn.Where(&database.Choice{User: user.ID, Choice: req.Choice}).Take(&choice).Error != nil {
		return util.FailedRequest(c, "Du hast dieses Spiel nicht ausgewählt!", nil)
	}

	// Remove choice
	if err := database.DBConn.Delete(&choice).Error; err != nil {
		return util.FailedRequest(c, "Beim Löschen ist ein unerwarteter Fehler aufgetreten. Bitte versuche es später nochmal!", nil)
	}

	return util.SuccessfulRequest(c)
}
