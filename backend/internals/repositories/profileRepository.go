package repositories

import (
	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db}
}

func (r *ProfileRepository) Create(profile *models.Profile) (*models.Profile, error) {
	err := r.db.Create(profile).Error
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (r *ProfileRepository) FindById(id uint) (*models.Profile, error) {
	var profile models.Profile
	err := r.db.First(&profile, id).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepository) FindByUserID(userID uint) (*models.Profile, error) {
	var profile models.Profile
	err := r.db.Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepository) Update(oldProfile *models.Profile, profile *models.Profile) (*models.Profile, error) {
	if profile.FirstName != "" && profile.FirstName != oldProfile.FirstName {
		oldProfile.FirstName = profile.FirstName
	}
	if profile.LastName != "" && profile.LastName != oldProfile.LastName {
		oldProfile.LastName = profile.LastName
	}
	if !profile.BirthDate.IsZero() && profile.BirthDate != oldProfile.BirthDate {
		oldProfile.BirthDate = profile.BirthDate
	}
	if profile.Avatar != "" && profile.Avatar != oldProfile.Avatar {
		oldProfile.Avatar = profile.Avatar
	}
	if profile.Bio != "" && profile.Bio != oldProfile.Bio {
		oldProfile.Bio = profile.Bio
	}

	err := r.db.Save(oldProfile).Error
	if err != nil {
		return nil, err
	}
	return oldProfile, nil
}

func (r *ProfileRepository) Delete(id uint) error {
	err := r.db.Delete(&models.Profile{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
