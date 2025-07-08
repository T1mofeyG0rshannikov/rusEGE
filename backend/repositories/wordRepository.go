package repositories

import (
	"rusEGE/database/models"
	"rusEGE/exceptions"

	"gorm.io/gorm"
)

type WordRepository interface {
	CreateUserWord(word *models.UserWord) error
	Get(id *uint) (*models.Word, error)
	Edit(word *models.Word) (*models.Word, error)
	GetAll() ([]*models.Word, error)
	GetRule(rule string) (*models.Rule, error)
	CreateRule(rule *models.Rule) (*models.Rule, error)
	Delete(word string) error
	CreateError(userId, wordId uint) error
	GetWordErrors(wordId uint) (*[]models.Error, error)
	GetRuleErrorsCount(ruleId uint) (*int64, error)
}

type GormWordRepository struct {
	db *gorm.DB
}

func NewGormWordRepository(db *gorm.DB) *GormWordRepository {
	return &GormWordRepository{db: db}
}

func (r *GormWordRepository) GetWordErrors(userId, wordId uint) (*[]models.Error, error) {
	var errors *[]models.Error
	result := r.db.Where("word_id = ? AND user_id = ?", wordId, userId).Find(&errors)
	if result.Error != nil {
		return nil, result.Error
	}

	return errors, nil
}

func (r *GormWordRepository) CreateUserWord(word *models.UserWord) (*models.UserWord, error) {
	result := r.db.Create(word)
	if result.Error != nil {
		return nil, result.Error
	}

	return word, nil
}

func (r *GormWordRepository) CreateError(userId, wordId uint) (*models.Error, error) {
	wordError := &models.Error{
		UserId: userId,
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
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *GormWordRepository) Create(word *models.Word) (*models.Word, error) {
	result := r.db.Create(word)
	if result.Error != nil {
		return nil, result.Error
	}

	return word, nil
}

func (r *GormWordRepository) Get(id *uint) (*models.Word, error) {
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

func (r *GormWordRepository) CreateRule(rule *models.Rule) (*models.Rule, error) {
	result := r.db.Create(rule)
	if result.Error != nil {
		return nil, result.Error
	}

	return rule, nil
}

func (r *GormWordRepository) GetAll() ([]*models.Word, error) {
	var words []*models.Word
	result := r.db.Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}

func (r *GormWordRepository) GetRule(ruleContent string) (*models.Rule, error) {
	var rule *models.Rule
	result := r.db.Where("Rule = ?", ruleContent).First(&rule)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, exceptions.ErrRuleNotFound
		}

		return nil, result.Error
	}

	return rule, nil
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

		result := r.db.Preload("Rule").Where("task_id = ? AND rule_id IN (?)", taskId, interfaceSlice).Find(&words)
		if result.Error != nil {
			return nil, result.Error
		}

		return words, nil
	} else {
		result := r.db.Preload("Rule").Where("task_id = ?", taskId).Find(&words)
		if result.Error != nil {
			return nil, result.Error
		}

		return words, nil
	}
}

func (r *GormWordRepository) GetTaskUserWords(taskId, userId uint) ([]*models.UserWord, error) {
	var words []*models.UserWord
	result := r.db.Where("TaskId = ? AND UserId = ?", taskId, userId).Find(&words)
	if result.Error != nil {
		return nil, result.Error
	}

	return words, nil
}

func (r *GormWordRepository) GetRuleErrorsCount(ruleId, userId uint) (*int64, error) {
	var count int64
	result := r.db.Table("rules").
		Select("count(errors.id)").
		Joins("INNER JOIN words ON rules.id = words.rule_id").
		Joins("INNER JOIN errors ON words.id = errors.word_id").
		Where("rules.id = ? AND errors.user_id = ?", ruleId, userId). // Добавлено условие по UserId
		Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	return &count, nil
}
