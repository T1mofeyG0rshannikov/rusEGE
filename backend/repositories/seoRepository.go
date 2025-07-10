package repositories

import (
	"rusEGE/database"
	"rusEGE/database/models"
	"rusEGE/exceptions"

	"gorm.io/gorm"
)

type SeoRepository interface {
	CreateIndexSeo(word *models.IndexSeo) error
	GetIndexSeo() (*models.IndexSeo, error)
}

type GormSeoRepository struct {
	db *gorm.DB
}

func NewGormSeoRepository(db *gorm.DB) *GormSeoRepository {
	if db == nil{
		db = database.GetDB()
	}
	return &GormSeoRepository{db}
}

func (r *GormSeoRepository) GetIndexSeo() (*models.IndexSeo, error) {
	var seo *models.IndexSeo
	result := r.db.First(&seo)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, exceptions.ErrIndexSeoNotFound
		}

		return nil, result.Error
	}

	return seo, nil
}

func (r *GormSeoRepository) CreateIndexSeo(seo *models.IndexSeo) (*models.IndexSeo, error) {
	result := r.db.Create(seo)
	if result.Error != nil {
		return nil, result.Error
	}

	return seo, nil
}


func (r *GormSeoRepository) EditIndexSeo(seo *models.IndexSeo) (*models.IndexSeo, error) {
	result := r.db.Save(&seo)
	if result.Error != nil {
		return nil, result.Error
	}
	return seo, nil
}
