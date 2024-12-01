package validation

import (
	"errors"

	"github.com/iacopoghilardi/mydget-backend/internals/types/dto"
	"github.com/iacopoghilardi/mydget-backend/utils"
)

func ValidateLoginUserDto(dto *dto.LoginUserDto) error {
	if dto.Email == "" || dto.Password == "" {
		return errors.New("email and password are required")
	}

	if !utils.IsEmailValid(dto.Email) {
		return errors.New("invalid email")
	}

	return nil
}

func ValidateRegisterUserDto(dto *dto.RegisterUserDto) error {
	if dto.Email == "" || dto.Password == "" {
		return errors.New("email and password are required")
	}

	if !utils.IsEmailValid(dto.Email) {
		return errors.New("invalid email")
	}

	passwordValidator := NewPasswordValidator()

	if err := passwordValidator.Validate(dto.Password); err != nil {
		return err
	}

	if err := passwordValidator.ValidateMatch(dto.Password, dto.ConfirmPassword); err != nil {
		return err
	}

	return nil
}

func ValidateCreateUserDto(dto *dto.CreateUserDto) error {
	if dto.Email == "" || dto.Password == "" {
		return errors.New("email and password are required")
	}

	if !utils.IsEmailValid(dto.Email) {
		return errors.New("invalid email")
	}

	passwordValidator := NewPasswordValidator()

	if err := passwordValidator.Validate(dto.Password); err != nil {
		return err
	}

	return nil
}

func ValidateUpdateUserDto(dto *dto.UpdateUserDto) error {
	return nil
}