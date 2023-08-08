package caching

import (
	"time"

	"github.com/Unbreathable/sportfest/backend/util"
	"github.com/dgraph-io/ristretto"
)

// ! Always use cost 1
var emailCache *ristretto.Cache // Email -> Token
const TokenTTL = time.Hour * 1  // 1 hour

func SetupCaches() {
	var err error
	emailCache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e5, // number of objects expected (100,000).
		MaxCost:     1e5, // maximum cost of cache (100,000).
		BufferItems: 64,  // something from the docs
	})
	if err != nil {
		panic(err)
	}
}

type EmailToken struct {
	Verified bool   // Whether the email has been verified
	Attempts int    // Number of attempts to verify the email
	Code     string // Code to verify email
	Year     uint   // Year of the account
	Password string // Password of the account
}

// Returns true if the email is in the cache
func AddEmailToVerify(email string) (EmailToken, bool) {

	// Check if email is already in cache
	_, valid := emailCache.Get(email)
	if valid {
		return EmailToken{}, false
	}

	token := EmailToken{
		Verified: false,
		Attempts: 0,
		Code:     util.GenerateToken(6),
		Year:     0,
		Password: "",
	}
	emailCache.SetWithTTL(email, token, 1, TokenTTL)

	return token, true
}

// Removes the email from the cache
func RemoveEmail(email string) {
	emailCache.Del(email)
}

// Returns true if the email is verified (code can be empty)
func VerifyEmail(email string, code string) (EmailToken, bool) {

	token, valid := emailCache.Get(email)
	if !valid {
		return EmailToken{}, false
	}

	emailToken := token.(EmailToken)
	if emailToken.Verified {
		return emailToken, true
	}

	if emailToken.Code == code {
		emailToken.Verified = true
		emailCache.SetWithTTL(email, emailToken, 1, TokenTTL)
		return emailToken, true
	}

	emailToken.Attempts++
	emailCache.SetWithTTL(email, emailToken, 1, TokenTTL)
	return emailToken, false
}

func SetPassword(email string, password string) {
	token, valid := emailCache.Get(email)
	if !valid {
		return
	}

	emailToken := token.(EmailToken)
	emailToken.Password = password
	emailCache.SetWithTTL(email, emailToken, 1, TokenTTL)
}

func SetYear(email string, year uint) {
	token, valid := emailCache.Get(email)
	if !valid {
		return
	}

	emailToken := token.(EmailToken)
	emailToken.Year = year
	emailCache.SetWithTTL(email, emailToken, 1, TokenTTL)
}
