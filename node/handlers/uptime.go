package handlers

import (
	internal "iskra/node/internal/uptime"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUptime(c echo.Context) error {
	uptime := internal.Uptime()
	return c.JSON(http.StatusOK, uptime)
}
