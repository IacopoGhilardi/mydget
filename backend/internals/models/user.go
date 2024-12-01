package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string   `gorm:"size:255; column:first_name" json:"first_name"`
	LastName  string   `gorm:"size:255; column:last_name" json:"last_name"`
	Email     string   `gorm:"uniqueIndex;size:255; column:email" json:"email"`
	Password  string   `gorm:"size:255; column:password" json:"password"`
	Profile   *Profile `gorm:"foreignKey:UserID" json:"profile"`
}

func (u *User) Create(db *gorm.DB) error {
	if err := u.validate(); err != nil {
		return err
	}
	return db.Create(u).Error
}

func (u *User) Update(db *gorm.DB) error {
	if err := u.validate(); err != nil {
		return err
	}
	return db.Save(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}

func (u *User) validate() error {
	if u.FirstName == "" || u.LastName == "" {
		return errors.New("first name and last name cannot be empty")
	}
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}
	return nil
}
