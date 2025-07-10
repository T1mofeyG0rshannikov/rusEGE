package handlers

import (
	"net/http"
	"rusEGE/database"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
	"rusEGE/web/utils"

	"github.com/labstack/echo/v4"

	usecases "rusEGE/usecases/rules"
	"strconv"
)


func EditRuleHadler(c echo.Context) error {
	var payload schemas.EditRuleRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	db := database.GetDB()

	rr := repositories.NewGormRuleRepository(db)

	rule, err := usecases.EditRule(rr, payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"rule": rule,
	})
}

func GetRulesStatHandler(c echo.Context) error {
	numberStr := c.Param("task")
	number, err := strconv.ParseUint(numberStr, 10, 64)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	taskNumber := uint(number)
	user := utils.UserFromContext(c)

	db := database.GetDB()
	tr := repositories.NewGormTaskRepository(db)
	rr := repositories.NewGormRuleRepository(db)

	stat, err := usecases.GetRulesStat(tr, rr, taskNumber, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"stat": stat,
	})
}
