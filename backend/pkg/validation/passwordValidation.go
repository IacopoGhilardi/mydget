package validation

import (
	"errors"
	"regexp"
)

var (
	minLength      = 8
	requireNumber  = true
	requireUpper   = true
	requireLower   = true
	requireSpecial = true

	ErrPasswordTooShort  = errors.New("password must be at least 8 characters long")
	ErrPasswordNoNumber  = errors.New("password must contain at least one number")
	ErrPasswordNoUpper   = errors.New("password must contain at least one uppercase letter")
	ErrPasswordNoLower   = errors.New("password must contain at least one lowercase letter")
	ErrPasswordNoSpecial = errors.New("password must contain at least one special character")
	ErrPasswordsNotMatch = errors.New("passwords do not match")
)

type PasswordValidator struct {
	MinLength      int
	RequireNumber  bool
	RequireUpper   bool
	RequireLower   bool
	RequireSpecial bool
}

func NewPasswordValidator() *PasswordValidator {
	return &PasswordValidator{
		MinLength:      minLength,
		RequireNumber:  requireNumber,
		RequireUpper:   requireUpper,
		RequireLower:   requireLower,
		RequireSpecial: requireSpecial,
	}
}

func (v *PasswordValidator) Validate(password string) error {
	if len(password) < v.MinLength {
		return ErrPasswordTooShort
	}

	if v.RequireNumber && !regexp.MustCompile(`\d`).MatchString(password) {
		return ErrPasswordNoNumber
	}

	if v.RequireUpper && !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return ErrPasswordNoUpper
	}

	if v.RequireLower && !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return ErrPasswordNoLower
	}

	if v.RequireSpecial && !regexp.MustCompile(`[@$!%*#?&]`).MatchString(password) {
		return ErrPasswordNoSpecial
	}

	return nil
}

func (v *PasswordValidator) ValidateMatch(password, confirmPassword string) error {
	if password != confirmPassword {
		return ErrPasswordsNotMatch
	}
	return nil
}
