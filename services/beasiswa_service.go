package services

import (
	"errors"
	"time"

	"github.com/farhandwiyan/project-himaose/models"
	"github.com/farhandwiyan/project-himaose/repositories"
	"github.com/google/uuid"
)

type BeasiswaService interface {
	Create(besiswa *models.Beasiswa) error
	Update(beasiswa *models.Beasiswa) error
	GetByPublicID(publicID string) (*models.Beasiswa, error)
	GetAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.Beasiswa, int64, error)
	DeleteBeasiswaByID(publicID string) error
	GetAllByStatusOprec() ([]models.Beasiswa, error)
}

type beasiswaService struct {
	beasiswaRepo repositories.BeasiswaRepository
	userRepo repositories.UserRepository
}

func NewBeasiswaService(beasiswaRepo repositories.BeasiswaRepository, userRepo repositories.UserRepository) BeasiswaService {
	return &beasiswaService{beasiswaRepo, userRepo}
}

func (s *beasiswaService) Create(beasiswa *models.Beasiswa) error {
	if beasiswa.CreatedBy == 0 {
		return errors.New("user id is required")
	}

	// cek user
	user, err := s.userRepo.FindByID(beasiswa.CreatedBy)
	if err != nil {
		return errors.New("owner not found")
	}

	beasiswa.PublicID = uuid.New()
	beasiswa.CreatedBy = user.InternalID
	return s.beasiswaRepo.Create(beasiswa)
}

func (s *beasiswaService) Update(beasiswa *models.Beasiswa) error {
	return s.beasiswaRepo.Update(beasiswa)
}

func (s *beasiswaService) GetByPublicID(publicID string) (*models.Beasiswa, error) {
	return s.beasiswaRepo.FindByPublicID(publicID)
}

func (s *beasiswaService) GetAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.Beasiswa, int64, error) {
	return s.beasiswaRepo.FindAllByUserPaginate(userID, filter, sort, limit, offset)
}

func (s *beasiswaService) DeleteBeasiswaByID(publicID string) error {
	return s.beasiswaRepo.Delete(publicID)
}

func (s *beasiswaService) GetAllByStatusOprec() ([]models.Beasiswa, error) {
	now := time.Now().Format("2006-01-02")
	
	return s.beasiswaRepo.FindByStatus(now)
}