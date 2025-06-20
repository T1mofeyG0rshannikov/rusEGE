package main


import (
	"rusEGE/web/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("user/get", handlers.GetUserHandler)
	e.POST("/login", handlers.LoginHandler)
	e.POST("/register", handlers.CreateUserHandler)

	e.Start(":8080")
}
