package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: change to secured key
const secretkey = "supersecret"

// function for generation of tokens
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretkey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected sigining method")
		}
		return []byte(secretkey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	
	// Extract userId as float64, then convert
	uidFloat, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("userId claim is missing or not a number")
	}

	return int64(uidFloat), nil
}
