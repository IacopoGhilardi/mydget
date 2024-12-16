package services

import (
	"errors"

	"github.com/iacopoghilardi/mydget-backend/internals/mappers"
	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"github.com/iacopoghilardi/mydget-backend/internals/repositories"
	"github.com/iacopoghilardi/mydget-backend/internals/types/dto"
	"github.com/iacopoghilardi/mydget-backend/utils"
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

func (a *AuthService) Login(user *dto.LoginUserDto) (string, error) {
	userModel := mappers.LoginUserDtoToUserModel(user)

	existingUser, err := a.userRepository.FindByEmail(userModel.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utils.VerifyPassword(existingUser.Password, userModel.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(existingUser.UUID)
	if err != nil {
		return "", errors.New("failed to generate JWT")
	}

	return token, nil
}

func (a *AuthService) ResetPassword(user *models.User) (*models.User, error) {
	return nil, nil
}
