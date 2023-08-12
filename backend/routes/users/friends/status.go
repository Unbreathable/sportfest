package friends

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type statusRequest struct {
	Code string `json:"code"`
}

// Route: /users/friends/status
func status(c *fiber.Ctx) error {

	// Parse request
	var req statusRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	var user database.User
	if database.DBConn.Where(&database.User{Code: req.Code}).Take(&user).Error != nil {
		return util.InvalidRequest(c)
	}

	// Get friendship status
	var friendship database.Friendship
	if database.DBConn.Where(&database.Friendship{ID: user.ID}).Take(&friendship).Error != nil {

		// No friendship found
		return c.JSON(fiber.Map{
			"success":     true,
			"code":        "",
			"set":         false,
			"friends":     false,
			"enteredCode": "",
		})
	}

	if friendship.FriendCode == "" {

		// Friend code found, but no friends yet
		return c.JSON(fiber.Map{
			"success":     true,
			"code":        friendship.Code,
			"set":         true,
			"friends":     false,
			"enteredCode": "",
		})
	}

	// Check if friendship is mutual
	var friend database.Friendship
	if database.DBConn.Where(&database.Friendship{Code: friendship.FriendCode}).Take(&friend).Error != nil {

		// Friend code found, but mutual
		return c.JSON(fiber.Map{
			"success":     true,
			"code":        friendship.Code,
			"set":         true,
			"friends":     false,
			"enteredCode": friendship.FriendCode,
		})
	}

	if friend.Code == friendship.FriendCode && friend.FriendCode == friendship.Code {

		// Get user
		var friendUser database.User
		if database.DBConn.Where(&database.User{ID: friend.ID}).Take(&friendUser).Error != nil {
			return util.FailedRequest(c, "Eure Freundschaft scheint registriert zu sein, aber dein Freund existiert nicht mehr in der Datenbank. Falls du diesen Fehler siehst, benachrichtige bitte einen Administrator.", nil)
		}

		// Friendship is mutual
		return c.JSON(fiber.Map{
			"success":     true,
			"code":        friendship.Code,
			"set":         true,
			"friends":     true,
			"friend":      friendUser.Name,
			"enteredCode": friendship.FriendCode,
		})
	}

	// Friendship is not mutual
	return c.JSON(fiber.Map{
		"success":     true,
		"code":        friendship.Code,
		"set":         true,
		"friends":     false,
		"enteredCode": friendship.FriendCode,
	})
}
