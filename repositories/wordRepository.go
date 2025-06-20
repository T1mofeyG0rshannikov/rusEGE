package repositories

import (
	"rusEGE/database"
	"gorm.io/gorm"
)


type WordRepository interface {
	CreateUserWord(word *database.UserWord) error
	GetAll() ([] *database.Word, error)
}

type GormWordRepository struct {
	db *gorm.DB
}

func NewGormWordRepository(db *gorm.DB) *GormWordRepository {
	return &GormWordRepository{db: db}
}

func (r *GormWordRepository) CreateUserWord(word *database.UserWord) (*database.UserWord, error) {
	result := r.db.Create(word)
	if result.Error != nil{
		return nil, result.Error
	}

	return word, nil
}

func (r *GormWordRepository) GetAll() ([] *database.Word, error) {
	var words []*database.Word
	result := r.db.Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}