package middleware

import (
	"net/http"
	"rusEGE/auth"
	"rusEGE/database"
	"rusEGE/repositories"
	"rusEGE/web/config"
	"rusEGE/web/utils"

	"github.com/labstack/echo/v4"
)

func UserRequiredMiddleware() echo.MiddlewareFunc {
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

			c.Set(config.UserContextKey, user)

			return next(c)
		}
	}
}
