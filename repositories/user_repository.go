package repositories

import (
	"github.com/farhandwiyan/project-himaose/config"
	"github.com/farhandwiyan/project-himaose/models"
)

type UserRepository interface {
	// berisi function repository yang dimiliki user
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByID(id int64) (*models.User, error)
}

// tidak dapat diakses diluar package agar menyembunyikan bisnis logic
type userRepository struct {

}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	
	return &user, err
}

func (r *userRepository) FindByID(id int64) (*models.User, error) {
	var user models.User
	err := config.DB.Where("internal_id = ?", id).First(&user).Error
	return &user, err
}

