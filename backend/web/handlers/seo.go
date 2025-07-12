package handlers

import (
	"errors"
	"net/http"
	"rusEGE/exceptions"
	"rusEGE/repositories"
	"rusEGE/usecases/seo"
	"rusEGE/web/schemas"

	"github.com/labstack/echo/v4"
)

func CreateIndexSeoHandler(c echo.Context) error {
	var payload schemas.CreateIndexSeoRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	sr := repositories.NewGormSeoRepository(nil)

	seo, err := seo.CreateIndexSeo(
		sr,
		payload,
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
	sr := repositories.NewGormSeoRepository(nil)

	seo, err := sr.GetIndexSeo()

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrIndexSeoNotFound):
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
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

	sr := repositories.NewGormSeoRepository(nil)

	seo, err := seo.EditIndexSeo(
		sr,
		payload,
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
