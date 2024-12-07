package handlers

import (
	internal "iskra/node/internal/cpuinfo"

	"github.com/labstack/echo/v4"
)

func GetCpuInfo(c echo.Context) error {
	cpuEntires := internal.CPUInfo()

	return c.JSON(200, cpuEntires)
}
