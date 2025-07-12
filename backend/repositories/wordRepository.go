package repositories

import (
	"errors"
	"rusEGE/database"
	"rusEGE/database/models"
	"rusEGE/exceptions"

	"gorm.io/gorm"
)

type WordRepository interface {
	CreateUserWord(word *models.UserWord) error
	Get(id *uint) (*models.Word, error)
	Edit(word *models.Word) (*models.Word, error)
	GetAll() ([]*models.Word, error)
	Delete(word string) error
	CreateError(userId, wordId uint) error
	GetWordErrors(wordId uint) (*[]models.Error, error)
}

type GormWordRepository struct {
	db *gorm.DB
}

func NewGormWordRepository(db *gorm.DB) *GormWordRepository {
	if db == nil {
		db = database.GetDB()
	}
	return &GormWordRepository{db}
}

func (r *GormWordRepository) DeleteOptions(wordId uint) error {
	result := r.db.Where("word_id = ?", wordId).Delete(&models.WordOption{})
	return result.Error
}

func (r *GormWordRepository) CreateOption(wordId uint, letter string) (*models.WordOption, error) {
	option := &models.WordOption{
		WordId: wordId,
		Letter: letter,
	}

	result := r.db.Create(&option)
	if result.Error != nil {
		return nil, result.Error
	}

	return option, nil
}

func (r *GormWordRepository) GetWithOptions(id uint) (*models.Word, error) {
	var word *models.Word
	result := r.db.Preload("Options").Where("Id = ?", id).First(&word)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, exceptions.ErrRuleNotFound
		}

		return nil, result.Error
	}

	return word, nil
}

func (r *GormWordRepository) CreateError(wordId uint) (*models.Error, error) {
	wordError := &models.Error{
		WordId: wordId,
	}

	result := r.db.Create(wordError)

	if result.Error != nil {
		return nil, result.Error
	}

	return wordError, nil
}

func (r *GormWordRepository) Delete(word string) error {
	var wordToDelete models.Word
	result := r.db.Where("word = ?", word).First(&wordToDelete)
	if result.Error != nil {
		return result.Error
	}

	result = r.db.Delete(wordToDelete)
	return result.Error
}

func (r *GormWordRepository) Get(id uint) (*models.Word, error) {
	var word *models.Word
	result := r.db.Where("Id = ?", id).First(&word)
	if result.Error != nil {
		return nil, result.Error
	}

	return word, nil
}

func (r *GormWordRepository) Edit(word *models.Word) (*models.Word, error) {
	result := r.db.Save(&word)
	if result.Error != nil {
		return nil, result.Error
	}
	return word, nil
}

func (r *GormWordRepository) All() ([]*models.Word, error) {
	var words []*models.Word
	result := r.db.Preload("Options").Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}

func contains(s []uint, str uint) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func (r *GormWordRepository) GetTaskWords(taskId uint, ruleIds *[]uint) ([]*models.Word, error) {
	var words []*models.Word
	if ruleIds != nil && len(*ruleIds) != 0 && !contains(*ruleIds, 0) {
		interfaceSlice := make([]interface{}, len(*ruleIds))
		for i, d := range *ruleIds {
			interfaceSlice[i] = d
		}

		result := r.db.Preload("Rule").Preload("Options").Preload("Rule.Options").Where("task_id = ? AND rule_id IN (?)", taskId, interfaceSlice).Find(&words)
		if result.Error != nil {
			return nil, result.Error
		}

		return words, nil
	} else {
		result := r.db.Preload("Rule").Preload("Options").Preload("Rule.Options").Where("task_id = ?", taskId).Find(&words)
		if result.Error != nil {
			return nil, result.Error
		}

		return words, nil
	}
}

func (r *GormWordRepository) GetByWord(wordContent string) (*models.Word, error) {
	var word *models.Word
	result := r.db.Where("word = ?", wordContent).First(&word)
	if result.Error != nil {
		return nil, result.Error
	}

	return word, nil
}

func (r *GormWordRepository) Create(wordContent string, taskId uint, ruleId uint, exception *bool, description *string) (*models.Word, error) {
	word := &models.Word{
		Word:   wordContent,
		TaskId: taskId,
		RuleId: ruleId,
		Exception: *exception,
		Description: description,
	}

	result := r.db.Create(word)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, exceptions.ErrWordAlreadyExists
		}
		return nil, result.Error
	}

	return word, nil
}