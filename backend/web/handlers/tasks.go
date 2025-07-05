package handlers

import (
	"errors"
	"net/http"
	"rusEGE/database"
	"rusEGE/exceptions"
	"rusEGE/interfaces"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
	"rusEGE/web/utils"

	"github.com/labstack/echo/v4"

	"math/rand"
	"rusEGE/auth"
	"rusEGE/usecases"
	"strconv"
	"time"
)

func shuffleSlice(slice []interfaces.Word) {
	// Инициализируем генератор случайных чисел.
	rand.Seed(time.Now().UnixNano())

	// Реализация перемешивания Фишера-Йетса (Fisher-Yates shuffle).
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func CreateTaskHandler(c echo.Context) error {
	var payload schemas.CreateTaskRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	db := database.GetDB()
	tr := repositories.NewGormTaskRepository(db)

	task, err := tr.Create(&database.Task{
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

	task, err := tr.Get(payload.TaskNumber)
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

	word, err := wr.Create(&database.Word{
		TaskId: task.Id,
		Word:   payload.Word,
		Rule:   payload.Rule,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"word": word,
	})
}

func GetWordsHandler(c echo.Context) error {
	db := database.GetDB()

	user, err := utils.GetUserFromHeader(
		auth.NewJWTProcessor(),
		repositories.NewGormUserRepository(db),
		c,
	)

	numberStr := c.Param("taskNumber")
	number, err := strconv.ParseUint(numberStr, 10, 64)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	numberUint := uint(number)

	words, err := usecases.GetTaskWords(numberUint, user)

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

	shuffleSlice(words)
	data := map[string]interface{}{
		"words": words,
	}

	return c.JSON(http.StatusOK, data)
}


func GetTasksHandler(c echo.Context) error {
	db := database.GetDB()
	tr := repositories.NewGormTaskRepository(db)

	tasks, err := tr.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
}
