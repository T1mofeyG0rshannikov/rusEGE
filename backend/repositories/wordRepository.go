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

func (r *GormWordRepository) Create(word *database.Word) (*database.Word, error) {
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


func (r *GormWordRepository) GetTaskWords(taskId uint) ([] *database.Word, error) {
	var words []*database.Word
	result := r.db.Where("task_id = ?", taskId).Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}


func (r *GormWordRepository) GetTaskUserWords(taskId, userId uint) ([] *database.UserWord, error) {
	var words []*database.UserWord
	result := r.db.Where("TaskId = ? UserId = ?", taskId, userId).Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}
