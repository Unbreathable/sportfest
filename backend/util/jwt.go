package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Tokens have no expiration in this project since it isn't a vulnerable application in general,
// and we don't want to have to deal with refreshing tokens

// JWT_SECRET is the secret used to sign the jwt token
func Token(account uint) (string, error) {

	// Create jwt token
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"acc": account,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := tk.SignedString([]byte(JWT_SECRET))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetAcc(c *fiber.Ctx) uint {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return uint(claims["acc"].(float64))
}
