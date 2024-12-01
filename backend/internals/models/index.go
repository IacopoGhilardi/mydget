package models

import "gorm.io/gorm"

type Model interface {
	Create(db *gorm.DB) error
	Update(db *gorm.DB) error
	Delete(db *gorm.DB) error
}
