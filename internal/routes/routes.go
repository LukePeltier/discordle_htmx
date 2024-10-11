package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/lukepeltier/discordle/internal/handlers"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.HomeHandler)
	e.GET("/play", handlers.PlayHandler)
	e.GET("/api/randomMessage", handlers.RandomMessageAPIHandler)
	e.POST("/api/messages/:id", handlers.MessageGuessAPIHandler)
}
