package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	Email   string    `json:"email"`
	User_ID uuid.UUID `json:"userid"`
	jwt.RegisteredClaims
}

func VerifyToken(tokenString string) (*CustomClaims, error) {
	secretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token or claims")
	}

	return claims, nil
}
