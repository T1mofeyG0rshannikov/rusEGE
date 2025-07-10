package handlers

import (
	"errors"
	"net/http"
	"rusEGE/database"
	"rusEGE/exceptions"
	"rusEGE/repositories"
	usecases "rusEGE/usecases/words"
	"rusEGE/web/schemas"
	"rusEGE/web/utils"

	"github.com/labstack/echo/v4"
)

func CreateWordHandler(c echo.Context) error {
	var payload schemas.CreateWordRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := database.GetDB()
	tr := repositories.NewGormTaskRepository(db)
	wr := repositories.NewGormWordRepository(db)
	rr := repositories.NewGormRuleRepository(db)

	word, err := usecases.CreateWord(tr, wr, rr, payload)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrTaskNotFound):
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "something went wrong",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"word": word,
	})
}

func BulkCreateWordHandler(c echo.Context) error {
	var payload schemas.BulkCreateWordRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := database.GetDB()
	tr := repositories.NewGormTaskRepository(db)
	wr := repositories.NewGormWordRepository(db)
	rr := repositories.NewGormRuleRepository(db)

	err := usecases.BulkCreateWord(tr, wr, rr, payload)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrTaskNotFound):
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "something went wrong",
			})
		}
	}

	return c.String(http.StatusOK, "")
}

func EditWordHandler(c echo.Context) error {
	var payload schemas.EditWordRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := database.GetDB()
	wr := repositories.NewGormWordRepository(db)
	rr := repositories.NewGormRuleRepository(db)

	word, err := usecases.EditWord(wr, rr, payload)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrTaskNotFound):
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "something went wrong",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"word": word,
	})
}

func GetWordsHandler(c echo.Context) error {
	var payload schemas.GetWordsRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	db := database.GetDB()

	user := utils.UserFromContext(c)

	tr := repositories.NewGormTaskRepository(db)
	wr := repositories.NewGormWordRepository(db)

	words, err := usecases.GetTaskWords(tr, wr, payload, user)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrTaskNotFound):
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "something went wrong",
			})
		}
	}

	data := map[string]interface{}{
		"words": words,
	}

	return c.JSON(http.StatusOK, data)
}

func DeleteWordHandler(c echo.Context) error {
	var payload schemas.DeleteWordsRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	db := database.GetDB()
	wr := repositories.NewGormWordRepository(db)
	err := wr.Delete(payload.Word)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.String(http.StatusOK, "")
}

func CreateWordErrorHandler(c echo.Context) error {
	var payload schemas.CreateWordErrorRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	db := database.GetDB()

	user := utils.UserFromContext(c)
	wr := repositories.NewGormWordRepository(db)

	userError, wordError, err := usecases.CreateError(wr, payload, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"user_error": map[string]interface{}{
				"word_id": userError.WordId,
			},
			"word_error": map[string]interface{}{
				"word_id": wordError.WordId,
			},
		},
	})
}

func DeleteUserErrorHandler(c echo.Context) error {
	var payload schemas.DeleteUserErrorRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user := utils.UserFromContext(c)
	
	db := database.GetDB()
	wr := repositories.NewGormWordRepository(db)

	err := wr.DeleteUserError(payload.Word, user.Id)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrRecordNotFound):
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
	}

	return c.String(http.StatusOK, "")
}
