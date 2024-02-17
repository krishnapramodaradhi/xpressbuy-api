package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func FetchToken(payload string) (string, error) {
	claims := struct {
		Payload string `json:"payload"`
		jwt.RegisteredClaims
	}{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "xpressbuy",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SIGNING_KEY")))
}
