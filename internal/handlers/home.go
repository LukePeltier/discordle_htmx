package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/lukepeltier/discordle/internal/templates/home"
)

type HomeData struct {
}

func HomeHandler(c echo.Context) error {
	return render(c, home.Home())
}
