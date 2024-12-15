package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func ComparePassword(password, newPassword string) bool {
	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword)) == nil
}
