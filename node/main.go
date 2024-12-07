package main

import (
	"iskra/node/handlers"
	"log"
	"runtime"

	"github.com/labstack/echo/v4"
)

func main() {
	if runtime.GOOS == "windows" {
		log.Fatalf("This program does not support Windows currently.")
	}

	e := echo.New()
	registerRoutes(e)
	e.Start(":8080")
}

func registerRoutes(e *echo.Echo) {
	apiRoutes := e.Group("/api/v1")
	apiRoutes.GET("/meminfo", handlers.GetMemInfo)
	apiRoutes.GET("/cpuinfo", handlers.GetCpuInfo)
	apiRoutes.GET("/uptime", handlers.GetUptime)
}
