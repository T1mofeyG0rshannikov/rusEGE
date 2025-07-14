package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func IndexPageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"contentTemplate": "index",
	})
}

func TaskPageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "task.html", map[string]interface{}{
		"contentTemplate": "task",
	})
}

func TasksPageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "tasks.html", map[string]interface{}{
		"contentTemplate": "tasks",
	})
}

func StatisticsPageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "statistics.html", map[string]interface{}{
		"contentTemplate": "statistics",
	})
}
