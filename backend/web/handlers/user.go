package handlers

import (
	"errors"
	"net/http"
	"rusEGE/auth"
	"rusEGE/database"
	"rusEGE/exceptions"
	"rusEGE/repositories"
	"rusEGE/security"
	"rusEGE/usecases"
	"rusEGE/web/schemas"
	"rusEGE/web/utils"

	"github.com/labstack/echo/v4"
)

func CreateUserHandler(c echo.Context) error {
	db := database.GetDB()

	var payload schemas.CreateUserRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	accessToken, err := usecases.CreateUser(
		repositories.NewGormUserRepository(db),
		repositories.NewGormWordRepository(db),
		auth.NewJWTProcessor(),
		security.NewScryptHasher(),
		&payload,
	)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrUsernameExist):
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "something went wrong",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": accessToken,
	})
}

func LoginHandler(c echo.Context) error {
	db := database.GetDB()

	var payload schemas.LoginRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	accessToken, err := usecases.LoginUser(
		repositories.NewGormUserRepository(db),
		security.NewScryptHasher(),
		auth.NewJWTProcessor(),
		&payload,
	)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrUserNotFound):
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": err.Error(),
			})

		case errors.Is(err, exceptions.ErrIncorrectPassword):
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "something went wrong",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": accessToken,
	})
}

func GetUserHandler(c echo.Context) error {
	db := database.GetDB()

	user, err := utils.GetUserFromHeader(
		auth.NewJWTProcessor(),
		repositories.NewGormUserRepository(db),
		c,
	)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrUserNotFound):
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": err.Error(),
			})

		case errors.Is(err, exceptions.ErrNoAuthHeader):
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		case errors.Is(err, exceptions.ErrInvalidJwtToken):
			return c.JSON(http.StatusNonAuthoritativeInfo, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "something went wrong",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"username": user.Username,
		"user_id":  user.Id,
	})
}
