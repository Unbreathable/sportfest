package account

import (
	"github.com/Unbreathable/sportfest/backend/caching"
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type finishYearRequest struct {
	Email string `json:"email"`
	Year  uint   `json:"year"`
}

// Handles setting the year for a user currently registering
func registerYear(c *fiber.Ctx) error {

	// Parse request
	var req finishYearRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Normalize email
	email := NormalizeEmail(req.Email)

	// Get email token
	_, valid := caching.VerifyEmail(email, "")
	if !valid {
		return util.FailedRequest(c, "Deine Email ist noch nicht verifiziert!", nil)
	}

	// Set and verify year
	var year database.Year
	if err := database.DBConn.Where(&database.Year{ID: req.Year}).Take(&year).Error; err != nil {
		return util.FailedRequest(c, "Dieser Jahrgang existiert nicht mehr!", nil)
	}
	caching.SetYear(email, year.ID)

	return util.SuccessfulRequest(c)
}

type yearListRequest struct {
	Email string `json:"email"`
}

// Handles getting the list of years for a user currently registering
func registerYearList(c *fiber.Ctx) error {

	// Parse request
	var req yearListRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Normalize email
	email := NormalizeEmail(req.Email)

	// Get email token
	_, valid := caching.VerifyEmail(email, "")
	if !valid {
		return util.FailedRequest(c, "Deine Email ist noch nicht verifiziert!", nil)
	}

	// Get years
	var years []database.Year
	if err := database.DBConn.Find(&years).Error; err != nil {
		return util.FailedRequest(c, "Es ist ein Fehler aufgetreten!", nil)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"obj":     years,
	})
}
