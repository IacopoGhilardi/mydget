package mappers

import (
	"log"

	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"github.com/iacopoghilardi/mydget-backend/internals/types/dto"
	"github.com/iacopoghilardi/mydget-backend/utils"
)

func GetUserDtoFromCreateUserDto(dto *dto.CreateUserDto) models.User {
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	return models.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  hashedPassword,
	}
}

func RegisterUserDtoToUserModel(dto *dto.RegisterUserDto) models.User {
	return GetUserDtoFromCreateUserDto(&dto.CreateUserDto)
}
