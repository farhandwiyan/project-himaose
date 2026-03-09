package repositories

import (
	"errors"

	"github.com/farhandwiyan/project-himaose/config"
	"github.com/farhandwiyan/project-himaose/models"
)

type ProgramKerjaRepository interface {
	Create(proker *models.ProgramKerja) error
	Update(proker *models.ProgramKerja) error  
	FindByPublicID(PublicID string) (*models.ProgramKerja, error)
	FindAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.ProgramKerja, int64, error) 
	Delete(publicID string) error
	FindByStatus(status string) ([]models.ProgramKerja, error)
}

type programKerjaRepository struct {
}

func NewProgramKerjaRepository() ProgramKerjaRepository {
	return &programKerjaRepository{}
}

func (r *programKerjaRepository) Create(proker *models.ProgramKerja) error {
	return config.DB.Create(proker).Error
}

func (r *programKerjaRepository) Update(proker *models.ProgramKerja) error {
	return config.DB.Model(&models.ProgramKerja{}).
		Where("public_id = ?", proker.PublicID).Updates(map[string]interface{}{
		"nama_proker": proker.NamaProker,
		"deskripsi": proker.Deskripsi,
		"divisi": proker.Divisi,
		"status": proker.Status,
		"link_oprec": proker.LinkOprec,
	}).Error
}

func (r *programKerjaRepository) FindByPublicID(PublicID string) (*models.ProgramKerja, error) {
	var proker models.ProgramKerja
	err := config.DB.Where("public_id = ?", PublicID).First(&proker).Error
	return &proker, err
}

func (r *programKerjaRepository) FindAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.ProgramKerja, int64, error) {
	var proker []models.ProgramKerja
	var total int64
	
	query := config.DB.Model(&models.ProgramKerja{}).Where("created_by = ?", userID)
	
	if filter != "" {
		query = query.Where("nama_proker LIKE ?", "%"+filter+"%")
	}

	// count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// sorting created at
	if sort != "" {
		query = query.Order(sort)
	} else {
		query = query.Order("created_at desc")
	}

	if err := query.Limit(limit).Offset(offset).Find(&proker).Error; err != nil{
		return nil, 0, err
	}

	return proker, total, nil
}

func (r *programKerjaRepository) Delete(publicID string) error {
    result := config.DB.Where("public_id = ?", publicID).Delete(&models.ProgramKerja{})
    
    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return errors.New("no record found with that public_id")
    }

    return nil
}

func (r *programKerjaRepository) FindByStatus(status string) ([]models.ProgramKerja, error) {
	var proker []models.ProgramKerja

	query := config.DB.Model(&models.ProgramKerja{}).Where("status = ?", status)

	if query.Error != nil {
		return nil, errors.New("no record found with that status")
	}

	if err := query.Find(&proker).Error; err != nil {
		return nil, err 
	}

	return proker, nil
}