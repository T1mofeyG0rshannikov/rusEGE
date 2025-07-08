package middleware

import (
	"net/http"
	"rusEGE/auth"
	"rusEGE/database"
	"rusEGE/repositories"
	"rusEGE/web/utils"

	"github.com/labstack/echo/v4"
)

// Ключ контекста для хранения информации о пользователе
const userContextKey = "user"

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db := database.GetDB()

			user, err := utils.GetUserFromHeader(
				auth.NewJWTProcessor(),
				repositories.NewGormUserRepository(db),
				c,
			)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized: " + err.Error()})
			}

			// Создаем новый контекст с информацией о пользователе
			c.Set(userContextKey, user)

			return next(c)
		}
	}
}
