package main

import (
	"fmt"
	"iskra/centralized/internal/database"
	"iskra/centralized/internal/handlers"
	"iskra/centralized/internal/middlewares"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(c echo.Context) error {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return err
	}
	defer conn.Close()

	fmt.Println("Client connected", conn.RemoteAddr())

	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Printf("Received: %s\n", message)

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}

	return nil
}

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
		AllowOrigins:     []string{"*"},
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

	// Testing
	e.GET("/ws", func(c echo.Context) error {
		return handleWebSocket(c)
	})
}
