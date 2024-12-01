package services

import (
	"github.com/iacopoghilardi/mydget-backend/internals/mappers"
	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"github.com/iacopoghilardi/mydget-backend/internals/repositories"
	"github.com/iacopoghilardi/mydget-backend/internals/types/dto"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{userRepository}
}

func (a *AuthService) Register(user *dto.RegisterUserDto) (*models.User, error) {
	userModel := mappers.RegisterUserDtoToUserModel(user)

	createdUser, err := a.userRepository.CreateUser(&userModel)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (a *AuthService) Login(user *models.User) (*models.User, error) {
	return nil, nil
}

func (a *AuthService) ResetPassword(user *models.User) (*models.User, error) {
	return nil, nil
}
