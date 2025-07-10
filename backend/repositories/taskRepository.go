package repositories

import (
	"errors"
	"rusEGE/database"
	"rusEGE/database/models"
	"rusEGE/exceptions"
	"rusEGE/web/schemas"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(word *models.UserWord) error
	GetAll() ([]*models.Word, error)
	Get(number uint) (*models.Task, error)
	GetTaskRules(id uint) ([]*models.Rule, error)
}

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) *GormTaskRepository {
	if db == nil{
		db = database.GetDB()
	}
	return &GormTaskRepository{db}
}

func (r *GormTaskRepository) GetAll() ([]*models.Task, error) {
	var tasks []*models.Task
	result := r.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (r *GormTaskRepository) GetTaskRules(id uint) ([]*models.Rule, error) {
	var rules []*models.Rule
	result := r.db.Where("Id IN (?)", r.db.Model(&models.Word{}).Select("RuleId").Where("task_id = ?", id)).Find(&rules)
	if result.Error != nil {
		return nil, result.Error
	}

	return rules, nil
}

func (r *GormTaskRepository) Create(task *models.Task) (*models.Task, error) {
	result := r.db.Create(task)
	if result.Error != nil {
		switch {
		case errors.Is(result.Error, gorm.ErrDuplicatedKey):
			return nil, exceptions.ErrTaskAlreadyExists
		default:
			return nil, result.Error
		}
	}

	return task, nil
}

func (r *GormTaskRepository) Edit(number uint, data schemas.EditTaskRequest) error {
	task, err := r.Get(number)
	if err != nil {
		return err
	}

	task.Description = data.Description

	result := r.db.Save(&task)
	if result.Error != nil {
		return err
	}
	return nil
}

func (r *GormTaskRepository) Get(number uint) (*models.Task, error) {
	var task models.Task
	result := r.db.Where("Number = ?", number).First(&task)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrTaskNotFound
		} else {
			return nil, result.Error
		}
	}
	return &task, nil
}
