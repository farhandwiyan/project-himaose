package repositories

import (
	"errors"

	"github.com/farhandwiyan/project-himaose/config"
	"github.com/farhandwiyan/project-himaose/models"
)

type LombaRepository interface {
	Create(lomba *models.Lomba) error
	Update(lomba *models.Lomba) error
	FindByPublicID(publicID string) (*models.Lomba, error)
	FindAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.Lomba, int64, error)
	Delete(publicID string) error
	FindByStatus(closeDate string) ([]models.Lomba, error)
}

type lombaRepository struct {

}

func NewLombaRepository() LombaRepository {
	return &lombaRepository{}
}

func (r *lombaRepository) Create(lomba *models.Lomba) error {
	return config.DB.Create(lomba).Error
}

func (r *lombaRepository) Update(lomba *models.Lomba) error {
	return config.DB.Model(&models.Lomba{}).
	Where("public_id = ?", lomba.PublicID).Updates(map[string]interface{}{
		"nama_lomba": lomba.NamaLomba,
		"deskripsi_lomba":lomba.DeskripsiLomba,
		"persyaratan": lomba.Persyaratan,
		"tgl_buka": lomba.TglBuka,
		"tgl_tutup": lomba.TglTutup,
	}).Error
}

func (r *lombaRepository) FindByPublicID(publicID string) (*models.Lomba, error) {
	var lomba models.Lomba
	err := config.DB.Where("public_id = ?", publicID).First(&lomba).Error
	return &lomba, err
}

func (r *lombaRepository) FindAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.Lomba, int64, error) {
	var lomba []models.Lomba
	var total int64

	query := config.DB.Model(&models.Lomba{}).Where("created_by = ?", userID)

	if filter != "" {
		query = query.Where("nama_lomba LIKE ?", "%"+filter+"%")
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

	if err := query.Limit(limit).Offset(offset).Find(&lomba).Error; err != nil{
		return nil, 0, err
	}

	return lomba, total, nil
}

func (r *lombaRepository) Delete(publicID string) error {
    result := config.DB.Where("public_id = ?", publicID).Delete(&models.Lomba{})
    
    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return errors.New("no record found with that public_id")
    }

    return nil
}

func (r *lombaRepository) FindByStatus(closeDate string) ([]models.Lomba, error) {
	var lomba []models.Lomba

	query := config.DB.Model(models.Lomba{}).Where("tgl_tutup >= ?", closeDate)

	if query.Error != nil {
		return nil, errors.New("no record found with that status")
	}

	if err := query.Find(&lomba).Error; err != nil {
		return nil, err
	}

	return lomba, nil
}