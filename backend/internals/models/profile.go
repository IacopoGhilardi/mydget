package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID    int       `gorm:"column:user_id;uniqueIndex" json:"user_id"`
	FirstName string    `gorm:"column:first_name" json:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name"`
	BirthDate time.Time `gorm:"column:birth_date" json:"birth_date"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Bio       string    `gorm:"column:bio" json:"bio"`
}
