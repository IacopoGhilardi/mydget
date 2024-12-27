package dto

import "time"

type ProfileDto struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Avatar    string    `json:"avatar"`
	Bio       string    `json:"bio"`
}

type CreateProfileDto struct {
	ProfileDto
	UserID uint `json:"user_id" binding:"required"`
}

type UpdateProfileDto struct {
	ProfileDto
	UserID uint `json:"user_id" binding:"required"`
	ID     uint `json:"id" binding:"required"`
}
