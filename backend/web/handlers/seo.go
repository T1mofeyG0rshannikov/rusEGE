package handlers

import (
	"errors"
	"net/http"
	"rusEGE/auth"
	"rusEGE/database"
	"rusEGE/exceptions"
	"rusEGE/repositories"
	"rusEGE/usecases/seo"
	"rusEGE/web/schemas"
	"rusEGE/web/utils"

	"github.com/labstack/echo/v4"
)

func CreateIndexSeoHandler(c echo.Context) error {
	var payload schemas.CreateIndexSeoRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := database.GetDB()

	user, err := utils.GetUserFromHeader(
		auth.NewJWTProcessor(),
		repositories.NewGormUserRepository(db),
		c,
	)

	sr := repositories.NewGormSeoRepository(db)

	seo, err := seo.CreateIndexSeo(
		sr,
		payload,
		user,
	)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrIndexSeoAlreadyExists):
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"seo": seo,
	})
}

func GetIndexSeoHandler(c echo.Context) error {
	db := database.GetDB()

	sr := repositories.NewGormSeoRepository(db)

	seo, err := sr.GetIndexSeo()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"seo": seo,
	})
}

func EditIndexSeoHandler(c echo.Context) error {
	var payload schemas.EditIndexSeoRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := database.GetDB()

	user, err := utils.GetUserFromHeader(
		auth.NewJWTProcessor(),
		repositories.NewGormUserRepository(db),
		c,
	)

	sr := repositories.NewGormSeoRepository(db)

	seo, err := seo.EditIndexSeo(
		sr,
		payload,
		user,
	)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrIndexSeoAlreadyExists):
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"seo": seo,
	})
}
