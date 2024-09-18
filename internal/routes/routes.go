package routes

import (
	"github.com/labstack/echo/v5"
	"github.com/lukepeltier/discordle/internal/handlers"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.HomeHandler)
}
