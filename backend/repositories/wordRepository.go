package repositories

import (
	"rusEGE/database/models"
	"gorm.io/gorm"
)


type WordRepository interface {
	CreateUserWord(word *models.UserWord) error
	GetAll() ([] *models.Word, error)
}

type GormWordRepository struct {
	db *gorm.DB
}

func NewGormWordRepository(db *gorm.DB) *GormWordRepository {
	return &GormWordRepository{db: db}
}

func (r *GormWordRepository) CreateUserWord(word *models.UserWord) (*models.UserWord, error) {
	result := r.db.Create(word)
	if result.Error != nil{
		return nil, result.Error
	}

	return word, nil
}

func (r *GormWordRepository) Create(word *models.Word) (*models.Word, error) {
	result := r.db.Create(word)
	if result.Error != nil{
		return nil, result.Error
	}

	return word, nil
}

func (r *GormWordRepository) GetAll() ([] *models.Word, error) {
	var words []*models.Word
	result := r.db.Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}


func (r *GormWordRepository) GetTaskWords(taskId uint) ([] *models.Word, error) {
	var words []*models.Word
	result := r.db.Where("task_id = ?", taskId).Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}


func (r *GormWordRepository) GetTaskUserWords(taskId, userId uint) ([] *models.UserWord, error) {
	var words []*models.UserWord
	result := r.db.Where("TaskId = ? UserId = ?", taskId, userId).Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}
