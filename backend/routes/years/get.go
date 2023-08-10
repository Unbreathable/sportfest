package years

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type getRequest struct {
	ID uint `json:"id"`
}

// Route: /arq/years/get
func getYear(c *fiber.Ctx) error {

	// Parse request
	var req getRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	if req.ID == 0 {
		return util.InvalidRequest(c)
	}

	var year database.Year
	if database.DBConn.Where("id = ?", req.ID).Take(&year).Error != nil {
		return util.FailedRequest(c, "Dieses Jahr existiert nicht!", nil)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"obj":     year,
	})
}
