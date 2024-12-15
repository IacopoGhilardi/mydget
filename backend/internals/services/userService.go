package services

import (
	"fmt"
	"log"

	"github.com/iacopoghilardi/mydget-backend/internals/db"
	"github.com/iacopoghilardi/mydget-backend/internals/mappers"
	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"github.com/iacopoghilardi/mydget-backend/internals/repositories"
	"github.com/iacopoghilardi/mydget-backend/internals/types/dto"
)

var userRepository = repositories.NewUserRepository(db.GetDB())

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

	fmt.Printf("userModel: %+v\n", userModel)

	err := db.GetDB().Create(&userModel).Error
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return models.User{}, err
	}

	fmt.Printf("userModel: %+v\n", userModel)

	return userModel, nil
}

func (s *UserService) Update(user dto.UpdateUserDto) (models.User, error) {
	userModel := mappers.UpdateUserDtoToUserModel(&user)

	oldUser, err := s.userRepository.FindById(int(user.ID))
	if err != nil {
		log.Printf("error finding user: %+v\n", err)
		return models.User{}, err
	}

	updatedUser, err := s.userRepository.Update(oldUser, &userModel)
	if err != nil {
		log.Printf("error updating user: %+v\n", err)
		return models.User{}, err
	}

	return *updatedUser, nil
}

func (s *UserService) Delete(id int) error {
	err := s.userRepository.Delete(id)
	if err != nil {
		log.Printf("error deleting user: %+v\n", err)
		return err
	}
	return nil
}
