package web

import (
	"rusEGE/web/handlers"
	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	t := NewTemplateRenderer()
	e.Renderer = t
	e.Static("/static", "../public/static")

	e.GET("/", handlers.IndexPageHandler)
	e.GET("/tasks", handlers.TasksPageHandler)
	e.GET("/task/:number", handlers.TaskPageHandler)

	e.GET("api/tasks/get", handlers.GetTasksHandler)
	e.POST("api/tasks/create", handlers.CreateTaskHandler)
	e.POST("api/tasks/:number/edit", handlers.EditTaskHandler)
	e.POST("api/words/create", handlers.CreateWordHandler)
	e.GET("api/words/get/:taskNumber", handlers.GetWordsHandler)
	e.GET("user/get", handlers.GetUserHandler)
	e.POST("/login", handlers.LoginHandler)
	e.POST("/register", handlers.CreateUserHandler)

	e.Start(":8080")
}
