package handlers

import (
	internal "iskra/node/internal/meminfo"

	"github.com/labstack/echo/v4"
)

func GetMemInfo(c echo.Context) error {
	memEntires := internal.Meminfo()
	return c.JSON(200, memEntires)
}
