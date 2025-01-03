package dto

import (
	"time"

	"github.com/iacopoghilardi/mydget-backend/internals/models"
)

type CreateUserDto struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type UpdateUserDto struct {
	ID        uint   `json:"id" binding:"required"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

type LoginUserDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserDto struct {
	CreateUserDto
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type LoginResponseDto struct {
	AuthToken
}

type RegisterResponseDto struct {
	LoginResponseDto
	User models.User `json:"user"`
}

type AuthToken struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
