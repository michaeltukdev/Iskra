package main

import (
	"encoding/json"
	"fmt"
	"iskra/node/handlers"
	cpuinfo "iskra/node/internal/cpuinfo"
	meminfo "iskra/node/internal/meminfo"
	uptime "iskra/node/internal/uptime"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func websocketClient() {
	// Connect to the WebSocket server
	serverAddr := "ws://host.docker.internal:81/ws"
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	// Handle interrupt signal to gracefully close the connection
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Send messages to the server
	go func() {
		for {
			// Fetch CPU information, memory information, and system uptime
			// Soon to be disk information
			cpuInfo := cpuinfo.CPUInfo()
			memInfo := meminfo.Meminfo()
			uptime := uptime.Uptime()

			// Marshal the data into JSON format
			data, err := json.Marshal(map[string]interface{}{
				"cpu":    cpuInfo,
				"memory": memInfo,
				"uptime": uptime,
			})
			if err != nil {
				log.Println("Error marshalling data:", err)
				return
			}

			// Send the data to the server
			err = conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println("Error writing message:", err)
				return
			}

			time.Sleep(10 * time.Second)
		}
	}()

	// Wait for interrupt signal to close the connection
	<-interrupt
	fmt.Println("Interrupt received, closing connection...")
}

// Main function starts the node and initializes the web server and WebSocket client.
func main() {
	// Check if the operating system is Windows and exit if it is.
	if runtime.GOOS == "windows" {
		log.Fatalf("This program does not support Windows currently.")
	}

	// Initialize the Echo web server.
	e := echo.New()

	// Register all HTTP and WebSocket routes.
	RegisterRoutes(e)

	// Start the WebSocket client in a separate goroutine.
	go websocketClient()

	// Start the web server on port 8081.
	fmt.Println("Server started on :8081")
	if err := e.Start(":8081"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// RegisterRoutes sets up all the HTTP routes and WebSocket endpoints for the node.
func RegisterRoutes(e *echo.Echo) {
	// API routes under /api/v1
	apiRoutes := e.Group("/api/v1")
	apiRoutes.GET("/meminfo", handlers.GetMemInfo) // GET memory information
	apiRoutes.GET("/cpuinfo", handlers.GetCpuInfo) // GET CPU information
	apiRoutes.GET("/uptime", handlers.GetUptime)   // GET system uptime

	// WebSocket endpoint for handling WebSocket connections
	e.GET("/ws", handlers.WebsocketHandler)
}
