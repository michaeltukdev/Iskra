package main

import (
	"fmt"
	"iskra/centralized/internal/database"
	"iskra/centralized/internal/handlers"
	"iskra/centralized/internal/middlewares"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// func handleWebSocket(c echo.Context) error {
// 	channel, err := websocket.Accept(c.Response(), c.Request(), nil)
// 	if err != nil {
// 		log.Printf("Failed to accept WebSocket: %v\n", err)
// 		return err
// 	}
// 	defer channel.Close(websocket.StatusNormalClosure, "Normal closure")

// 	for {
// 		var v interface{}
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()

// 		err := wsjson.Read(ctx, channel, &v)
// 		if err != nil {
// 			switch websocket.CloseStatus(err) {
// 			case websocket.StatusNormalClosure, websocket.StatusGoingAway:
// 				return nil
// 			}

// 			log.Printf("Failed to read WebSocket message: %v\n", err)
// 			break
// 		}

// 		log.Printf("Received: %v\n", v)
// 	}

// 	return nil
// }

type Config struct {
	FRONTEND_URL string
	BACKEND_URL  string
	JWT_SECRET   string
}

func LoadConfig() *Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	backendURL := os.Getenv("BACKEND_URL")
	if strings.HasPrefix(backendURL, "http://") || strings.HasPrefix(backendURL, "https://") {
		backendURL = strings.Split(backendURL, "://")[1]
	}

	config := &Config{
		FRONTEND_URL: os.Getenv("FRONTEND_URL"),
		BACKEND_URL:  backendURL,
		JWT_SECRET:   os.Getenv("JWT_SECRET"),
	}

	if config.FRONTEND_URL == "" || config.BACKEND_URL == "" || config.JWT_SECRET == "" {
		log.Fatal("Missing required configuration")
	}

	return config
}

func main() {
	config := LoadConfig()
	fmt.Println("Starting server...")

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.FRONTEND_URL},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	RegisterRoutes(e)

	database.Init()

	fmt.Println("Server started at", config.BACKEND_URL)
	e.Logger.Fatal(e.Start(config.BACKEND_URL))
}

func RegisterRoutes(e *echo.Echo) {
	config := LoadConfig()
	auth := e.Group("/auth")
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)
	auth.POST("/logout", handlers.Logout)

	e.POST("/me", handlers.Me, middlewares.JWTMiddleware(config.JWT_SECRET))

	protected := e.Group("/protected")
	protected.Use(middlewares.JWTMiddleware(config.JWT_SECRET))
	protected.GET("", handlers.Restricted)
}
