package services

import (
	"github.com/iacopoghilardi/mydget-backend/internals/db"
	"github.com/iacopoghilardi/mydget-backend/internals/mappers"
	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"github.com/iacopoghilardi/mydget-backend/internals/repositories"
	"github.com/iacopoghilardi/mydget-backend/internals/types/dto"
)

var databaseInstance = db.GetDB()

var userRepository = repositories.NewUserRepository(databaseInstance)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetAll() ([]models.User, error) {
	users, err := s.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetById(id int) (models.User, error) {
	user, err := s.userRepository.FindById(id)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (s *UserService) Create(user *dto.CreateUserDto) (models.User, error) {
	userModel := mappers.GetUserDtoFromCreateUserDto(user)

	err := databaseInstance.Create(&userModel).Error
	if err != nil {
		return models.User{}, err
	}

	return userModel, nil
}

func (s *UserService) Update(user models.User) (models.User, error) {
	return user, nil
}

func (s *UserService) Delete(id int) error {
	return nil
}
