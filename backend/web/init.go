package web

import (
	"rusEGE/web/handlers"
	"rusEGE/web/middleware"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()

	t := NewTemplateRenderer()
	e.Renderer = t

	userRequiredGroup := e.Group("/")
	userRequiredGroup.Use(middleware.UserRequiredMiddleware())

	userRequiredGroup.POST("api/word-error/create", handlers.CreateWordErrorHandler)
	userRequiredGroup.DELETE("api/word-error/delete", handlers.DeleteUserErrorHandler)
	userRequiredGroup.GET("api/task/:number/stat", handlers.GetTaskStatHandler)
	userRequiredGroup.GET("api/rules/get-stat/:task", handlers.GetRulesStatHandler)
	userRequiredGroup.GET("api/user/get", handlers.GetUserHandler)

	userOptionalGroup := e.Group("/")
	userOptionalGroup.Use(middleware.OptionalUserMiddleware())
	userOptionalGroup.GET("api/words/get", handlers.GetWordsHandler)


	e.Static("/static", "../public/static")

	e.GET("/", handlers.IndexPageHandler)
	e.GET("/tasks", handlers.TasksPageHandler)
	e.GET("/task/:number", handlers.TaskPageHandler)

	e.POST("api/rule/edit", handlers.EditRuleHadler)
	e.GET("api/tasks/get", handlers.GetTasksHandler)
	e.POST("api/tasks/create", handlers.CreateTaskHandler)
	e.POST("api/tasks/:number/edit", handlers.EditTaskHandler)
	e.POST("api/words/create", handlers.CreateWordHandler)
	e.POST("api/words/bulk-create", handlers.BulkCreateWordHandler)
	e.POST("api/words/edit", handlers.EditWordHandler)
	e.DELETE("api/words/delete", handlers.DeleteWordHandler)
	e.POST("api/indexseo/create", handlers.CreateIndexSeoHandler)
	e.POST("api/indexseo/edit", handlers.EditIndexSeoHandler)
	e.GET("api/indexseo", handlers.GetIndexSeoHandler)
	e.POST("api/login", handlers.LoginHandler)
	e.POST("api/register", handlers.CreateUserHandler)
	e.POST("api/refresh-token/:token", handlers.RefreshTokenHandler)

	e.Start(":8080")
}
