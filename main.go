package main

import (
	"database/sql"
	"fmt"
	"iskra/centralized/internal/config"
	"iskra/centralized/internal/database"
	"iskra/centralized/internal/server"
	"log"
	"net/url"

	"github.com/labstack/echo/v4"
)

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func handleWebSocket(c echo.Context) error {
// 	// Upgrade the HTTP connection to a WebSocket connection
// 	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
// 	if err != nil {
// 		fmt.Println("Error upgrading connection:", err)
// 		return err
// 	}
// 	defer conn.Close()

// 	fmt.Println("Client connected", conn.RemoteAddr())

// 	for {
// 		// Read message from client
// 		messageType, message, err := conn.ReadMessage()
// 		if err != nil {
// 			fmt.Println("Error reading message:", err)
// 			break
// 		}

// 		fmt.Printf("Received: %s\n", message)

// 		err = conn.WriteMessage(messageType, message)
// 		if err != nil {
// 			fmt.Println("Error writing message:", err)
// 			break
// 		}
// 	}

// 	return nil
// }

type App struct {
	Config     *config.Config
	DB         *sql.DB
	HTTPClient *echo.Echo
}

func main() {
	config := config.Initialize()
	fmt.Println("Starting server...")

	db, err := database.Init()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	e := server.NewServer(config, db)
	parsedURL, err := url.Parse(config.BACKEND_URL)
	if err != nil {
		log.Fatalf("Error parsing backend URL: %v", err)
	}

	fmt.Println("Server started at", config.BACKEND_URL)
	e.Logger.Fatal(e.Start(parsedURL.Host))
}
