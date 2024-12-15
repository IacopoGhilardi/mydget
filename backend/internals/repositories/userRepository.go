package repositories

import (
	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"github.com/iacopoghilardi/mydget-backend/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindById(id int) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(oldUser *models.User, user *models.User) (*models.User, error) {
	if user.FirstName != "" && user.FirstName != oldUser.FirstName {
		oldUser.FirstName = user.FirstName
	}
	if user.LastName != "" && user.LastName != oldUser.LastName {
		oldUser.LastName = user.LastName
	}
	if user.Email != "" && user.Email != oldUser.Email {
		oldUser.Email = user.Email
	}
	if user.Password != "" && !utils.ComparePassword(user.Password, oldUser.Password) {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		oldUser.Password = hashedPassword
	}
	err := r.db.Save(oldUser).Error
	if err != nil {
		return nil, err
	}
	return oldUser, nil
}

func (r *UserRepository) Delete(id int) error {
	err := r.db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
