package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "superSecretKey"

func GenerateJWTTokenString(email string, userId int64) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})

	return jwtToken.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unsupported method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return -1, err
	}

	isValidToken := parsedToken.Valid
	if !isValidToken {
		return -1, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return -1, errors.New("Invalid claims type")
	}

	userId := int64(claims["user_id"].(float64))
	// _ = claims["email"]

	return userId, nil
}
