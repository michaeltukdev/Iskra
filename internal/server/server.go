package server

import (
	"iskra/centralized/internal/config"
	"iskra/centralized/internal/handlers"
	"iskra/centralized/internal/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

func NewServer(config *config.Config, db *bun.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.FRONTEND_URL},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	h := &handlers.Handlers{
		DB:        db,
		JWTSecret: config.JWT_SECRET,
	}

	auth := e.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
	auth.POST("/logout", h.Logout)

	// Protected route
	e.POST("/me", h.Me, middlewares.JWTMiddleware(config.JWT_SECRET))

	return e
}
