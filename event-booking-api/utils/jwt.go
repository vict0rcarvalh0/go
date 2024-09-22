package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "veryverysecretok?"

func GenerateToken(email string, userId int64) (string, error) {
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // the token will expire in 2 hours
	}) // create a new token with the HS256 algorithm

	return token.SignedString(secretKey)
}
