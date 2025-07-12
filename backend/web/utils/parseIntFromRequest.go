package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParseIntFromRequest(c echo.Context, key string) (*uint, error) {
	numberStr := c.Param(key)
	number, err := strconv.ParseUint(numberStr, 10, 64)

	if err != nil {
		return nil, err
	}

	numberUint := uint(number)
	return &numberUint, nil
}
