package repositories

import (
	"errors"
	"rusEGE/database"
	"rusEGE/database/models"
	"rusEGE/exceptions"

	"gorm.io/gorm"
)

type UserWordRepository interface {
	CreateUserWord(word *models.UserWord) error
	Get(id *uint) (*models.Word, error)
	Edit(word *models.Word) (*models.Word, error)
	GetAll() ([]*models.Word, error)
	Delete(word string) error
	CreateError(userId, wordId uint) error
	GetWordErrors(wordId uint) (*[]models.Error, error)
}

type GormUserWordRepository struct {
	db *gorm.DB
}

func NewGormUserWordRepository(db *gorm.DB) *GormUserWordRepository {
	if db == nil {
		db = database.GetDB()
	}
	return &GormUserWordRepository{db}
}

func (r *GormUserWordRepository) GetErrors(wordId uint) (*[]models.UserError, error) {
	var errors *[]models.UserError
	result := r.db.Where("word_id = ?", wordId).Find(&errors)
	if result.Error != nil {
		return nil, result.Error
	}

	return errors, nil
}

func (r *GormUserWordRepository) CreateOption(wordId uint, letter string) (*models.UserWordOption, error) {
	option := &models.UserWordOption{
		WordId: wordId,
		Letter: letter,
	}

	result := r.db.Create(&option)
	if result.Error != nil {
		return nil, result.Error
	}

	return option, nil
}

func (r *GormUserWordRepository) DeleteError(wordId uint) error {
	var userError models.UserError
	result := r.db.Where("word_id = ?", wordId).First(&userError)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return exceptions.ErrRecordNotFound
		} else {
			return result.Error
		}
	}

	result = r.db.Delete(&userError)
	return result.Error
}

func (r *GormUserWordRepository) Create(userId uint, wordContent string, taskId uint, ruleId uint, exception *bool, description *string) (*models.UserWord, error) {
	word := &models.UserWord{
		UserId:      userId,
		Word:        wordContent,
		TaskId:      taskId,
		RuleId:      ruleId,
		Exception:   *exception,
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

func (r *GormUserWordRepository) CreateError(userId, userWordId uint) (*models.UserError, error) {
	wordError := models.UserError{
		WordId: userWordId,
	}

	result := r.db.Create(&wordError)

	if result.Error != nil {
		return nil, result.Error
	}

	return &wordError, nil
}

func (r *GormUserWordRepository) GetTaskWords(taskId, userId uint, ruleIds *[]uint) ([]*models.UserWord, error) {
	var words []*models.UserWord
	if ruleIds != nil && len(*ruleIds) != 0 && !contains(*ruleIds, 0) {
		interfaceSlice := make([]interface{}, len(*ruleIds))
		for i, d := range *ruleIds {
			interfaceSlice[i] = d
		}

		result := r.db.Preload("Rule").Preload("Options").Preload("Rule.Options").Where("task_id = ? AND rule_id IN (?)", taskId, interfaceSlice).Find(&words)
		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		result := r.db.Preload("Rule").Preload("Options").Preload("Rule.Options").Where("task_id = ?", taskId).Find(&words)
		if result.Error != nil {
			return nil, result.Error
		}

	}
	return words, nil
}

func (r *GormUserWordRepository) Get(wordContent string) (*models.UserWord, error) {
	var word *models.UserWord
	result := r.db.Where("word = ?", wordContent).First(&word)
	if result.Error != nil {
		return nil, result.Error
	}

	return word, nil
}

func (r *GormUserWordRepository) Delete(word *models.UserWord) error {
	result := r.db.Delete(word)
	return result.Error
}

func (r *GormUserWordRepository) GetById(wordId uint) (*models.UserWord, error) {
	var word models.UserWord
	result := r.db.Where("id = ?", wordId).First(&word)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return nil, exceptions.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &word, nil
}
