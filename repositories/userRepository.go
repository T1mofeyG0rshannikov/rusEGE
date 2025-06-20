package repositories

import (
	"errors"
	"rusEGE/database"
	"rusEGE/exceptions"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *database.User) (*database.User, error)
	Get(username *string) (*database.User, error)
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user *database.User) (*database.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *GormUserRepository) Get(username string) (*database.User, error) {
	var user database.User
	result := r.db.Where("Username = ?", username).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, exceptions.ErrUserNotFound
		} else {
			return nil, result.Error
		}
	}
	return &user, nil
}
