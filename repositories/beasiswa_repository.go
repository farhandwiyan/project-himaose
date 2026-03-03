package repositories

import (
	"errors"

	"github.com/farhandwiyan/project-himaose/config"
	"github.com/farhandwiyan/project-himaose/models"
)

type BeasiswaRepository interface {
	Create(beasiswa *models.Beasiswa) error
	Update(beasiswa *models.Beasiswa) error
	FindByPublicID(publicID string) (*models.Beasiswa, error)
	FindAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.Beasiswa, int64, error) 
	Delete(publicID string) error
}

type beasiswaRepository struct {

}

func NewBeasiswaRepository() BeasiswaRepository {
	return &beasiswaRepository{}
}

func (r *beasiswaRepository) Create(beasiswa *models.Beasiswa) error {
	return config.DB.Create(beasiswa).Error
}

func (r *beasiswaRepository) Update(beasiswa *models.Beasiswa) error {
	return config.DB.Model(&models.Beasiswa{}).
	Where("public_id = ?", beasiswa.PublicID).Updates(map[string]interface{}{
		"nama_beasiswa":beasiswa.NamaBeasiswa,
		"link_pendaftaran":beasiswa.LinkPendaftaran,
		"tgl_buka":beasiswa.TglBuka,
		"tgl_tutup":beasiswa.TglTutup,
	}).Error
}	

func (r *beasiswaRepository) FindByPublicID(publicID string) (*models.Beasiswa, error) {
	var beasiswa models.Beasiswa
	err := config.DB.Where("public_id = ?", publicID).First(&beasiswa).Error
	return &beasiswa, err
}

func (r *beasiswaRepository) FindAllByUserPaginate(userID int64, filter,sort string, limit,offset int) ([]models.Beasiswa, int64, error) {
	var beasiswa []models.Beasiswa
	var total int64

	query := config.DB.Model(&models.Beasiswa{}).Where("created_by = ?", userID)

	if filter != "" {
		query = query.Where("nama_beasiswa LIKE ?", "%"+filter+"%")
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

	if err := query.Limit(limit).Offset(offset).Find(&beasiswa).Error; err != nil{
		return nil, 0, err
	}

	return beasiswa, total, nil
}

func (r *beasiswaRepository) Delete(publicID string) error {
    result := config.DB.Where("public_id = ?", publicID).Delete(&models.Beasiswa{})
    
    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return errors.New("no record found with that public_id")
    }

    return nil
}