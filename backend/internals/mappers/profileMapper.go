package mappers

import (
	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"github.com/iacopoghilardi/mydget-backend/internals/types/dto"
)

func GetProfileModelFromCreateProfileDto(dto *dto.CreateProfileDto) models.Profile {
	return models.Profile{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		BirthDate: dto.BirthDate,
		Avatar:    dto.Avatar,
		Bio:       dto.Bio,
		UserID:    dto.UserID,
	}
}
