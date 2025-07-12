package handlers

import (
	"errors"
	"net/http"
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

	tr := repositories.NewGormTaskRepository(nil)
	wr := repositories.NewGormWordRepository(nil)
	rr := repositories.NewGormRuleRepository(nil)

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

	tr := repositories.NewGormTaskRepository(nil)
	wr := repositories.NewGormWordRepository(nil)
	rr := repositories.NewGormRuleRepository(nil)

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

	wr := repositories.NewGormWordRepository(nil)
	rr := repositories.NewGormRuleRepository(nil)

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

	user := utils.UserFromContext(c)

	tr := repositories.NewGormTaskRepository(nil)
	wr := repositories.NewGormWordRepository(nil)
	uwr := repositories.NewGormUserWordRepository(nil)

	words, err := usecases.GetTaskWords(tr, wr, uwr, payload, user)

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

	wr := repositories.NewGormWordRepository(nil)
	err := wr.Delete(payload.Word)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.String(http.StatusOK, "")
}

func DeleteUserWordHandler(c echo.Context) error {
	wordId, err := utils.ParseIntFromRequest(c, "wordId")
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	user := utils.UserFromContext(c)
	uwr := repositories.NewGormUserWordRepository(nil)

	userWord, err := uwr.GetById(*wordId)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrRecordNotFound):
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "something went wrong",
			})
		}
	}

	if user.Id != userWord.UserId {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": "You don't have the rights to execute",
		})
	}

	err = uwr.Delete(userWord)

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

	user := utils.UserFromContext(c)
	wr := repositories.NewGormWordRepository(nil)
	uwr := repositories.NewGormUserWordRepository(nil)

	userError, wordError, err := usecases.CreateError(wr, uwr, payload, user)

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

func CreateUserWordHandler(c echo.Context) error {
	var payload schemas.CreateUserWordRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user := utils.UserFromContext(c)

	uwr := repositories.NewGormUserWordRepository(nil)
	tr := repositories.NewGormTaskRepository(nil)

	word, err := usecases.CreateUserWord(uwr, tr, payload, user)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrTaskNotFound):
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
		"word": word,
	})
}

func DeleteUserErrorHandler(c echo.Context) error {
	var payload schemas.DeleteUserErrorRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	wr := repositories.NewGormUserWordRepository(nil)

	err := wr.DeleteError(payload.Word)

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

func GetTaskUserWordsHandler(c echo.Context) error {
	taskNumber, err := utils.ParseIntFromRequest(c, "taskNumber")
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	user := utils.UserFromContext(c)

	uwr := repositories.NewGormUserWordRepository(nil)
	tr := repositories.NewGormTaskRepository(nil)

	words, err := usecases.GetTaskUserWords(
		uwr,
		tr,
		*taskNumber,
		user,
	)

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrTaskNotFound):
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
		"words": words,
	})
}

func GetWordErrorsHandler(c echo.Context) error {
	tr := repositories.NewGormTaskRepository(nil)
	wr := repositories.NewGormWordRepository(nil)

	stat, err := usecases.GetWordErrors(tr, wr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"stat": stat,
	})
}
