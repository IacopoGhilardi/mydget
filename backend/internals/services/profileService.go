package services

import (
	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"github.com/iacopoghilardi/mydget-backend/internals/types/dto"
	"gorm.io/gorm"
)

type ProfileService struct {
	db *gorm.DB
}

func NewProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{db}
}

func (s *ProfileService) GetProfile(id uint) (*models.Profile, error) {
	var profile models.Profile
	if err := s.db.Where("id = ?", id).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *ProfileService) CreateProfile(dto dto.CreateProfileDto) (*models.Profile, error) {
	profile := models.Profile{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		BirthDate: dto.BirthDate,
		Avatar:    dto.Avatar,
		Bio:       dto.Bio,
	}

	if err := s.db.Create(&profile).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}

func (s *ProfileService) UpdateProfile(dto dto.UpdateProfileDto) (*models.Profile, error) {
	var profile models.Profile
	if err := s.db.Where("id = ?", dto.ID).First(&profile).Error; err != nil {
		return nil, err
	}

	profile.FirstName = dto.FirstName
	profile.LastName = dto.LastName
	profile.BirthDate = dto.BirthDate
	profile.Avatar = dto.Avatar
	profile.Bio = dto.Bio

	if err := s.db.Save(&profile).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}
