// bisnis logic
package services

import (
	"errors"

	"github.com/farhandwiyan/project-himaose/models"
	"github.com/farhandwiyan/project-himaose/repositories"
	"github.com/farhandwiyan/project-himaose/utils"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *models.User) error
	Login(username,password string) (*models.User, error) 
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *models.User) error {
	// mengecek username yang terdaftar
	// hashing password
	// set role
	// simpan user

	existingUser, err := s.repo.FindByUsername(user.Username)
	if err == nil && existingUser != nil {
        return errors.New("username already registered")
    }

	hased, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hased
	user.Role = "admin"
	user.PublicID = uuid.New()

	return s.repo.Create(user)
}

func (s *userService) Login(username,password string) (*models.User, error) {
	user, err := s.repo.FindByUsername(username)

	if err != nil {
		return nil, errors.New("invalid credential")
	}
	
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credential")
	}

	return user, nil
}