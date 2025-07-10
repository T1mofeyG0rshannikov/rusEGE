package repositories

import (
	"errors"
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
	return &GormWordRepository{db: db}
}

func (r *GormWordRepository) GetUserWordErrors(wordId uint) (*[]models.UserError, error) {
	var errors *[]models.UserError
	result := r.db.Where("word_id = ?", wordId).Find(&errors)
	if result.Error != nil {
		return nil, result.Error
	}

	return errors, nil
}

func (r *GormWordRepository) DeleteUserError(wordId, userId uint) error {
	var userError models.UserError
	result := r.db.Where("word_id = ? AND user_id = ?", wordId, userId).First(&userError)
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

func (r *GormWordRepository) CreateUserWord(word *models.UserWord) (*models.UserWord, error) {
	result := r.db.Create(word)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey){
			return nil, exceptions.ErrWordAlreadyExists
		}
		return nil, result.Error
	}

	return word, nil
}

func (r *GormWordRepository) CreateUserError(userId, userWordId uint) (*models.UserError, error) {
	wordError := models.UserError{
		WordId: userWordId,
	}

	result := r.db.Create(&wordError)

	if result.Error != nil {
		return nil, result.Error
	}

	return &wordError, nil
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

func (r *GormWordRepository) Create(word *models.Word) (*models.Word, error) {
	result := r.db.Create(word)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey){
			return nil, exceptions.ErrWordAlreadyExists
		}
		return nil, result.Error
	}

	return word, nil
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

func (r *GormWordRepository) GetAll() ([]*models.Word, error) {
	var words []*models.Word
	result := r.db.Find(&words)
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

		result := r.db.Preload("Rule").Preload("Rule.Options").Where("task_id = ? AND rule_id IN (?)", taskId, interfaceSlice).Find(&words)
		if result.Error != nil {
			return nil, result.Error
		}

		return words, nil
	} else {
		result := r.db.Preload("Rule").Preload("Rule.Options").Where("task_id = ?", taskId).Find(&words)
		if result.Error != nil {
			return nil, result.Error
		}

		return words, nil
	}
}

func (r *GormWordRepository) GetTaskUserWords(taskId, userId uint, ruleIds *[]uint) ([]*models.UserWord, error) {
	var words []*models.UserWord
	if ruleIds != nil && len(*ruleIds) != 0 && !contains(*ruleIds, 0) {
		interfaceSlice := make([]interface{}, len(*ruleIds))
		for i, d := range *ruleIds {
			interfaceSlice[i] = d
		}

		result := r.db.Preload("Rule").Preload("Rule.Options").Where("task_id = ? AND rule_id IN (?)", taskId, interfaceSlice).Find(&words)
		if result.Error != nil {
			return nil, result.Error
		}

		return words, nil
	} else {
		result := r.db.Preload("Rule").Preload("Rule.Options").Where("task_id = ?", taskId).Find(&words)
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

func (r *GormWordRepository) GetUserWord(wordContent string) (*models.UserWord, error) {
	var word *models.UserWord
	result := r.db.Where("word = ?", wordContent).First(&word)
	if result.Error != nil {
		return nil, result.Error
	}

	return word, nil
}
