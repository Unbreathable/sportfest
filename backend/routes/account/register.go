package account

import (
	"regexp"
	"strings"

	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/gofiber/fiber/v2"
)

type registerRequest struct {
	Email string `json:"email"`
}

func registerAccount(c *fiber.Ctx) error {

	// Parse request
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		return util.InvalidRequest(c)
	}

	// Check if email is valid
	valid, email := CheckEmail(req.Email)
	if !valid {
		return util.InvalidRequest(c)
	}

	// Check if it's a school email

}

const EmailRegex = "^[a-zA-Z0-9]+(?:\\.[a-zA-Z0-9]+)*@[a-zA-Z0-9]+(?:\\.[a-zA-Z0-9]+)*$"

func NormalizeEmail(email string) string {

	// Convert email to lowercase
	email = strings.ToLower(email)

	// Remove leading and trailing whitespaces
	email = strings.TrimSpace(email)

	// Remove dots (.) from the username part of the email
	parts := strings.Split(email, "@")
	username := parts[0]
	username = strings.ReplaceAll(username, ".", "")

	// Reconstruct the normalized email address
	normalizedEmail := username + "@" + parts[1]

	return normalizedEmail
}

func CheckEmail(email string) (bool, string) {

	// Check if email is valid
	match, err := regexp.Match(EmailRegex, []byte(email))
	if !match || err != nil {
		return false, ""
	}

	email = NormalizeEmail(email)
	if strings.Contains(email, " ") {
		return false, ""
	}

	return true, email
}
