package utils

import (
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

func IsEmailValid(email string) bool {
	return emailRegex.MatchString(email)
}
