package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("mydget_secret_key")

type Claims struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}

func GenerateJWT(uuid string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UUID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Printf("Invalid token signature")
			return nil, errors.New("invalid token signature")
		}

		log.Printf("Error parsing JWT: %v", err)
		return nil, err
	}

	if !token.Valid {
		log.Printf("Invalid token")
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
