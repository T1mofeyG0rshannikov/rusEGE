package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexPageHandler(c echo.Context) error {
	data := map[string]interface{}{
		"contentTemplate": "index",
	}

	return c.Render(http.StatusOK, "index.html", data)
}

func TaskPageHandler(c echo.Context) error {
	data := map[string]interface{}{
		"contentTemplate": "task",
	}

	return c.Render(http.StatusOK, "task.html", data)
}

func TasksPageHandler(c echo.Context) error {
	data := map[string]interface{}{
		"contentTemplate": "tasks",
	}

	return c.Render(http.StatusOK, "tasks.html", data)
}
