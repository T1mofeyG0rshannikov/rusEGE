package repositories

import (
	"errors"
	"rusEGE/database/models"
	"rusEGE/exceptions"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	Get(username *string) (*models.User, error)
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db}
}

func (r *GormUserRepository) Create(username, hashedpassword string) (*models.User, error) {
	user := &models.User{
		Username:     username,
		HashPassword: hashedpassword,
	}

	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *GormUserRepository) Get(username string) (*models.User, error) {
	var user models.User
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
