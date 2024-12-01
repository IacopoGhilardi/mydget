package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID int `gorm:"column:user_id;uniqueIndex" json:"user_id"`
}
