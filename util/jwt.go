package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "superSecretKey"

func GenerateJWTTokenString(email string, userId int64) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 2),
	})

	return jwtToken.SignedString([]byte(secretKey))
}
