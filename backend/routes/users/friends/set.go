package friends

import (
	"strings"

	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type setRequest struct {
	Code       string `json:"code"`
	FriendCode string `json:"friend_code"`
}

// Route: /users/friends/set
func set(c *fiber.Ctx) error {

	// Parse request
	var req setRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	req.FriendCode = strings.TrimSpace(req.FriendCode)

	// Get user
	var user database.User
	if database.DBConn.Where(&database.User{Code: req.Code}).Take(&user).Error != nil {
		return util.InvalidRequest(c)
	}

	if req.FriendCode == "" {
		return util.FailedRequest(c, "Du musst einen Freundescode angeben.", nil)
	}

	var friendship database.Friendship
	if database.DBConn.Where(&database.Friendship{ID: user.ID}).Take(&friendship).Error != nil {
		return util.FailedRequest(c, "Du musst zuerst einen Freundescode generieren, bevor du einen Freund hinzufügen kannst.", nil)
	}

	if friendship.Code == req.FriendCode {
		return util.FailedRequest(c, "Du kannst dich nicht selbst hinzufügen.", nil)
	}

	friendship.FriendCode = req.FriendCode
	if database.DBConn.Save(&friendship).Error != nil {
		return util.FailedRequest(c, "Ein unerwarteter Fehler ist aufgetreten, bitte versuche es später erneut.", nil)
	}

	return util.SuccessfulRequest(c)
}
