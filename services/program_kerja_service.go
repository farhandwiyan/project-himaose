package services

import (
	"errors"

	"github.com/farhandwiyan/project-himaose/models"
	"github.com/farhandwiyan/project-himaose/repositories"
	"github.com/google/uuid"
)

type ProgramKerjaService interface {
	Create(proker *models.ProgramKerja) error 
	Update(proker *models.ProgramKerja) error
	GetByPublicID(publicID string) (*models.ProgramKerja, error)
	GetAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.ProgramKerja, int64, error) 
	DeleteProgramKerjaByID(publicID string) error 
}

type programKerjaService struct {
	programKerjaRepo repositories.ProgramKerjaRepository
	userRepo repositories.UserRepository
}

func NewProgramKerjaService(
	programKerjaRepo repositories.ProgramKerjaRepository, 
	userRepo repositories.UserRepository,
	) ProgramKerjaService {
	return &programKerjaService{programKerjaRepo, userRepo}
}

func (s *programKerjaService) Create(proker *models.ProgramKerja) error {
	if proker.CreatedBy == 0 {
        return errors.New("user id is required")
    }
	
	// cek user
	user, err := s.userRepo.FindByID(proker.CreatedBy)
	if err != nil {
		return errors.New("owner not found")
	}

	proker.PublicID = uuid.New()
	proker.CreatedBy = user.InternalID
	return s.programKerjaRepo.Create(proker)
}

func (s *programKerjaService) Update(proker *models.ProgramKerja) error {
	return s.programKerjaRepo.Update(proker)
}

func (s *programKerjaService) GetByPublicID(publicID string) (*models.ProgramKerja, error) {
	return s.programKerjaRepo.FindByPublicID(publicID)
}

func (s *programKerjaService) GetAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.ProgramKerja, int64, error) {
	return s.programKerjaRepo.FindAllByUserPaginate(userID, filter, sort, limit, offset)
}

func (s *programKerjaService) DeleteProgramKerjaByID(publicID string) error {
	return s.programKerjaRepo.Delete(publicID)
}