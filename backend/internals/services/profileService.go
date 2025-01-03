package services

import (
	"github.com/iacopoghilardi/mydget-backend/internals/mappers"
	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"github.com/iacopoghilardi/mydget-backend/internals/repositories"
	"github.com/iacopoghilardi/mydget-backend/internals/types/dto"
)

type ProfileService struct {
	profileRepository *repositories.ProfileRepository
}

func NewProfileService(profileRepository *repositories.ProfileRepository) *ProfileService {
	return &ProfileService{profileRepository: profileRepository}
}

func (s *ProfileService) GetProfile(id uint) (*models.Profile, error) {
	profile, err := s.profileRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *ProfileService) Create(dto dto.CreateProfileDto) (*models.Profile, error) {
	profile := mappers.GetProfileModelFromCreateProfileDto(&dto)

	createdProfile, err := s.profileRepository.Create(&profile)
	if err != nil {
		return nil, err
	}
	return createdProfile, nil
}

func (s *ProfileService) UpdateProfile(dto dto.UpdateProfileDto) (*models.Profile, error) {
	oldProfile, err := s.profileRepository.FindById(dto.ID)
	if err != nil {
		return nil, err
	}

	profile := models.Profile{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		BirthDate: dto.BirthDate,
		Avatar:    dto.Avatar,
		Bio:       dto.Bio,
	}

	updatedProfile, err := s.profileRepository.Update(oldProfile, &profile)
	if err != nil {
		return nil, err
	}

	return updatedProfile, nil
}
