package util

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Payload string `json:"payload"`
	jwt.RegisteredClaims
}

func GenerateToken(payload string) (string, error) {
	claims := CustomClaims{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "xpressbuy",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})

	switch {
	case token.Valid:
		if claims, ok := token.Claims.(*CustomClaims); ok {
			return claims.Payload, nil
		}
		return "", errors.New("unknown claim type")
	case errors.Is(err, jwt.ErrTokenMalformed):
		return "", errors.New("Invalid token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return "", errors.New("Invalid Signature")
	case errors.Is(err, jwt.ErrTokenExpired):
		return "", errors.New("Token Expired")
	default:
		log.Println(err)
		return "", errors.New("unknown issue occured while parsing the token")
	}
}
