package handlers

import (
	"errors"
	"net/http"
	"rusEGE/database"
	"rusEGE/database/models"
	"rusEGE/exceptions"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
	"rusEGE/web/utils"

	"github.com/labstack/echo/v4"

	"rusEGE/auth"
	"rusEGE/usecases"
	"strconv"
)

func CreateTaskHandler(c echo.Context) error {
	var payload schemas.CreateTaskRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := database.GetDB()
	tr := repositories.NewGormTaskRepository(db)

	task, err := tr.Create(&models.Task{
		Number: payload.Number,
	})

	if err != nil {
		switch {
		case errors.Is(err, exceptions.ErrTaskNotFound):
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": err.Error(),
			})
		case errors.Is(err, exceptions.ErrTaskAlreadyExists):
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
		"task": task,
	})
}

func EditTaskHandler(c echo.Context) error {
	var payload schemas.EditTaskRequest

	numberStr := c.Param("number")
	number, err := strconv.ParseUint(numberStr, 10, 64)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	numberUint := uint(number)

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := database.GetDB()
	tr := repositories.NewGormTaskRepository(db)

	err = tr.Edit(numberUint, payload)

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

func CreateWordHandler(c echo.Context) error {
	var payload schemas.CreateWordRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := database.GetDB()
	tr := repositories.NewGormTaskRepository(db)
	wr := repositories.NewGormWordRepository(db)

	word, err := usecases.CreateWord(tr, wr, payload)

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

	err := usecases.BulkCreateWord(tr, wr, payload)

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

	word, err := usecases.EditWord(wr, payload)

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

	user, err := utils.GetUserFromHeader(
		auth.NewJWTProcessor(),
		repositories.NewGormUserRepository(db),
		c,
	)

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

func GetTasksHandler(c echo.Context) error {
	db := database.GetDB()
	tr := repositories.NewGormTaskRepository(db)
	wr := repositories.NewGormWordRepository(db)

	user, err := utils.GetUserFromHeader(
		auth.NewJWTProcessor(),
		repositories.NewGormUserRepository(db),
		c,
	)

	tasks, err := usecases.GetTasks(tr, wr, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
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
	var payload schemas.CreateWordError

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	db := database.GetDB()

	user := utils.UserFromContext(c)
	if user != nil {
		wr := repositories.NewGormWordRepository(db)

		wordError, err := wr.CreateError(user.Id, payload.Word)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]uint{"word_error_id": wordError.Id, "word_id": wordError.WordId, "user_id": wordError.UserId})
	} else {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauth",
		})
	}
}

func GetTaskStatHandler(c echo.Context) error {
	numberStr := c.Param("number")
	number, err := strconv.ParseUint(numberStr, 10, 64)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	taskNumber := uint(number)

	db := database.GetDB()
	tr := repositories.NewGormTaskRepository(db)
	wr := repositories.NewGormWordRepository(db)

	user := utils.UserFromContext(c)

	stat, err := usecases.GetTaskStat(
		taskNumber,
		tr,
		wr,
		user,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"stat": stat,
	})
}
