package services

import (
	"errors"

	"github.com/farhandwiyan/project-himaose/models"
	"github.com/farhandwiyan/project-himaose/repositories"
	"github.com/google/uuid"
)

type LombaService interface {
	Create(lomba *models.Lomba) error
	Update(lomba *models.Lomba) error
	GetByPublicID(publicID string) (*models.Lomba, error) 
	GetAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.Lomba, int64, error)
	DeleteLombaByID(publicID string) error
}

type lombaService struct {
	lombaRepo repositories.LombaRepository
	userRepo repositories.UserRepository
}

func NewLombaService(
	lombaRepo repositories.LombaRepository, 
	userRepo repositories.UserRepository,
	) LombaService {
	return &lombaService{lombaRepo, userRepo}
}

func (s *lombaService) Create(lomba *models.Lomba) error {
	if lomba.CreatedBy == 0 {
		return errors.New("user id is required")
	}

	// cek user
	user, err := s.userRepo.FindByID(lomba.CreatedBy)
	if err != nil {
		return errors.New("owner not found")
	}

	lomba.PublicID = uuid.New()
	lomba.CreatedBy = user.InternalID
	return s.lombaRepo.Create(lomba)
}

func (s *lombaService) Update(lomba *models.Lomba) error {
	return s.lombaRepo.Update(lomba)
}

func (s *lombaService) GetByPublicID(publicID string) (*models.Lomba, error)  {
	return s.lombaRepo.FindByPublicID(publicID)
}

func (s *lombaService) GetAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.Lomba, int64, error) {
	return s.lombaRepo.FindAllByUserPaginate(userID, filter, sort, limit, offset)
}

func (s *lombaService) DeleteLombaByID(publicID string) error {
	return s.lombaRepo.Delete(publicID)
}