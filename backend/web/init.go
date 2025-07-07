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
	e.POST("api/words/bulk-create", handlers.BulkCreateWordHandler)
	e.POST("api/words/edit", handlers.EditWordHandler)
	e.DELETE("api/words/delete", handlers.DeleteWordHandler)
	e.GET("api/words/get", handlers.GetWordsHandler)
	e.POST("api/indexseo/create", handlers.CreateIndexSeoHandler)
	e.POST("api/indexseo/edit", handlers.EditIndexSeoHandler)
	e.GET("api/indexseo", handlers.GetIndexSeoHandler)
	e.GET("user/get", handlers.GetUserHandler)
	e.POST("/api/login", handlers.LoginHandler)
	e.POST("/api/register", handlers.CreateUserHandler)

	e.Start(":8080")
}
