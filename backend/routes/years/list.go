package years

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

// Route: /arq/years/list
func listYears(c *fiber.Ctx) error {

	var years []database.Year
	if database.DBConn.Find(&years).Error != nil {
		return util.FailedRequest(c, "Fehler beim Laden der Jahre!", nil)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"obj":     years,
	})
}
