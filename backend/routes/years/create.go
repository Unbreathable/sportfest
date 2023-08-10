package years

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type createYearRequest struct {
	Name string `json:"name"`
}

// Route: /arq/years/create
func createYear(c *fiber.Ctx) error {

	// Parse request
	var req createYearRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	if req.Name == "" {
		return util.FailedRequest(c, "Bitte gib einen Namen an!", nil)
	}

	if len(req.Name) < 4 {
		return util.FailedRequest(c, "Der Name muss mindestens 4 Zeichen lang sein!", nil)
	}

	var year database.Year
	if database.DBConn.Where("name = ?", req.Name).Take(&year).Error == nil {
		return util.FailedRequest(c, "Dieses Jahr existiert bereits!", nil)
	}

	// Create year
	year = database.Year{
		Name: req.Name,
	}
	database.DBConn.Create(&year)

	return c.JSON(fiber.Map{
		"success": true,
		"obj":     year.ID,
	})
}
