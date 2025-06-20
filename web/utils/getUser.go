package utils

import (
	"rusEGE/auth"
	"rusEGE/database"
	"rusEGE/exceptions"
	"rusEGE/repositories"

	"github.com/labstack/echo/v4"
)

func GetUserFromHeader(
	jwtProcessor *auth.JWTProcessor,
	ur *repositories.GormUserRepository,
	c echo.Context,
) (*database.User, error) {
	authHeader := c.Request().Header.Get("Authorization")
	
	if authHeader == ""{
		return nil, exceptions.ErrNoAuthHeader
	}

	claims, err := jwtProcessor.ValidateToken(authHeader)
	if err != nil {
		return nil, exceptions.ErrInvalidJwtToken
	}

	user, err := ur.Get(claims.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
