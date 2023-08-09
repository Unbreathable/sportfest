package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"math/big"

	"github.com/gofiber/fiber/v2"
)

var JWT_SECRET = "secret"
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateToken(tkLength int32) string {

	s := make([]rune, tkLength)

	length := big.NewInt(int64(len(letters)))

	for i := range s {

		number, _ := rand.Int(rand.Reader, length)
		s[i] = letters[number.Int64()]
	}

	return string(s)
}

func RandomInt(min, max int) int {
	number, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	return min + int(number.Int64())
}

func SuccessfulRequest(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}

func FailedRequest(c *fiber.Ctx, error string, err error) error {

	if err != nil {
		log.Println(c.Route().Path+":", err)
	}

	return c.Status(200).JSON(fiber.Map{
		"success": false,
		"error":   error,
	})
}

func InvalidRequest(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusBadRequest)
}

// HashPassword hashes a password
func HashPassword(password string) string {

	// Get hash of password (Should be secure enough)
	hash := sha256.Sum256([]byte(password))
	for i := 0; i < 50; i++ {
		hash = sha256.Sum256(hash[:])
	}

	return base64.StdEncoding.EncodeToString(hash[:])
}
