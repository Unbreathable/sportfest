package friends

import (
	"github.com/Unbreathable/sportfest/backend/database"
	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type friendsSimulationRequest struct {
	Friendships int `json:"friendships"`
}

// Route: /arq/users/friends/simulation
func friendsSimulation(c *fiber.Ctx) error {

	// Parse request
	var req friendsSimulationRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	var user []database.User
	if database.DBConn.Find(&user).Error != nil {
		return util.InvalidRequest(c)
	}

	var years []database.Year
	if database.DBConn.Find(&years).Error != nil {
		return util.InvalidRequest(c)
	}

	var alreadyDone []string
	for i := 0; i < req.Friendships; i++ {
		year := years[util.RandomInt(0, len(years))]

		// Befriend random users
		user1 := user[util.RandomInt(0, len(user))]
		user2 := user[util.RandomInt(0, len(user))]

		for user1.ID == user2.ID || contains(alreadyDone, user1.Name) || contains(alreadyDone, user2.Name) || user1.Year != year.ID || user2.Year != year.ID {
			user1 = user[util.RandomInt(0, len(user))]
			user2 = user[util.RandomInt(0, len(user))]
		}

		alreadyDone = append(alreadyDone, user1.Name)
		alreadyDone = append(alreadyDone, user2.Name)

		code1 := util.GenerateToken(12)
		code2 := util.GenerateToken(12)
		if err := database.DBConn.Create(&database.Friendship{
			ID:         user1.ID,
			Code:       code1,
			FriendCode: code2,
		}).Error; err != nil {
			return util.InvalidRequest(c)
		}
		if err := database.DBConn.Create(&database.Friendship{
			ID:         user2.ID,
			Code:       code2,
			FriendCode: code1,
		}).Error; err != nil {
			return util.InvalidRequest(c)
		}
	}

	return util.SuccessfulRequest(c)
}

func contains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
