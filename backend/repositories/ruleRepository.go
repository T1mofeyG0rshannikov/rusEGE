package repositories

import (
	"rusEGE/database/models"
	"rusEGE/exceptions"

	"gorm.io/gorm"
)

type RuleRepository interface {
	Get(rule string) (*models.Rule, error)
	GetById(id uint) (*models.Rule, error)
	GetWithOptions(id uint) (*models.Rule, error)
	CreateOption(models.RuleOption) (*models.RuleOption, error)
	DeleteOptions(ruleId uint) error
	Create(rule *models.Rule) (*models.Rule, error)
	GetTaskRules(id uint) ([]*models.Rule, error)
}

type GormRuleRepository struct {
	db *gorm.DB
}

func NewGormRuleRepository(db *gorm.DB) *GormRuleRepository {
	return &GormRuleRepository{db: db}
}

func (r *GormRuleRepository) Create(rule *models.Rule) (*models.Rule, error) {
	result := r.db.Create(rule)
	if result.Error != nil {
		return nil, result.Error
	}

	return rule, nil
}

func (r *GormRuleRepository) CreateOption(option *models.RuleOption) (*models.RuleOption, error) {
	result := r.db.Create(&option)
	if result.Error != nil {
		return nil, result.Error
	}

	return option, nil
}

func (r *GormRuleRepository) Get(ruleContent string) (*models.Rule, error) {
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

func (r *GormRuleRepository) GetWithOptions(id uint) (*models.Rule, error) {
	var rule *models.Rule
	result := r.db.Preload("Options").Where("Id = ?", id).First(&rule)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, exceptions.ErrRuleNotFound
		}

		return nil, result.Error
	}

	return rule, nil
}

func (r *GormRuleRepository) GetById(id uint) (*models.Rule, error) {
	var rule *models.Rule
	result := r.db.Where("Id = ?", id).First(&rule)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, exceptions.ErrRuleNotFound
		}

		return nil, result.Error
	}

	return rule, nil
}

func (r *GormRuleRepository) DeleteOptions(ruleId uint) error {
	result := r.db.Where("rule_id = ?", ruleId).Delete(&models.RuleOption{})
	return result.Error
}

func (r *GormRuleRepository) GetErrorsCount(ruleId, userId uint) (*int64, error) {
	var count int64
	result := r.db.Table("rules").
		Select("count(user_errors.id)").
		Joins("INNER JOIN user_words ON rules.id = user_words.rule_id").
		Joins("INNER JOIN user_errors ON user_words.id = user_errors.word_id").
		Where("rules.id = ? AND user_errors.user_id = ?", ruleId, userId).
		Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	return &count, nil
}

func (r *GormRuleRepository) GetTaskRules(id uint) ([]*models.Rule, error) {
	var rules []*models.Rule
	result := r.db.Where("Id IN (?)", r.db.Model(&models.Word{}).Select("RuleId").Where("task_id = ?", id)).Find(&rules)
	if result.Error != nil {
		return nil, result.Error
	}

	return rules, nil
}
