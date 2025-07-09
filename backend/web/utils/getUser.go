package utils

import (
	"rusEGE/auth"
	"rusEGE/database/models"
	"rusEGE/exceptions"
	"rusEGE/repositories"

	"github.com/labstack/echo/v4"
)

// Ключ контекста для хранения информации о пользователе
const userContextKey = "user"

func GetUserFromHeader(
	jwtProcessor *auth.JWTProcessor,
	ur *repositories.GormUserRepository,
	c echo.Context,
) (*models.User, error) {
	authHeader := c.Request().Header.Get("Authorization")

	if authHeader == "" {
		return nil, exceptions.ErrNoAuthHeader
	}

	claims, err := jwtProcessor.ValidateToken(authHeader)
	if err != nil {
		return nil, err
	}

	user, err := ur.Get(claims.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UserFromContext(c echo.Context) *models.User {
	user, ok := c.Get(userContextKey).(*models.User)
	if !ok {
		return nil // Или паникуйте, если это критическая ошибка
	}
	return user
}
