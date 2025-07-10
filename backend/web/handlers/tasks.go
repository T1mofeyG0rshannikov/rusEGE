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

	usecases "rusEGE/usecases/tasks"
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
		Number:      payload.Number,
		Description: payload.Description,
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

func GetTasksHandler(c echo.Context) error {
	db := database.GetDB()
	tr := repositories.NewGormTaskRepository(db)

	tasks, err := usecases.GetTasks(tr)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
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
