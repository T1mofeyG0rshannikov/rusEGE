package web

import (
	"log"
	"os"
	"rusEGE/web/handlers"
	"rusEGE/web/middleware"

	"github.com/joho/godotenv"
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
	userRequiredGroup.GET("api/user-words/get/:taskNumber", handlers.GetTaskUserWordsHandler)
	userRequiredGroup.DELETE("api/user-words/delete/:wordId", handlers.DeleteUserWordHandler)
	userRequiredGroup.POST("api/user-words/create", handlers.CreateUserWordHandler)

	userOptionalGroup := e.Group("/")
	userOptionalGroup.Use(middleware.OptionalUserMiddleware())
	userOptionalGroup.GET("api/words/get", handlers.GetWordsHandler)

	e.Static("/static", "../public/static")

	e.GET("/", handlers.IndexPageHandler)
	e.GET("/tasks", handlers.TasksPageHandler)
	e.GET("/task/:number", handlers.TaskPageHandler)
	e.GET("/statistics", handlers.StatisticsPageHandler)
	e.GET("/sitemap.xml", handlers.SitemapHandler)
	e.GET("api/word-errors/get", handlers.GetWordErrorsHandler)
	e.GET("api/rule/get/:taskNumber", handlers.GetTaskRulesHandler)
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

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла .env: %v", err)
	}

	port := os.Getenv("PORT")

	e.Start(":" + port)
}
