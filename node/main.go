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
	serverAddr := "ws://localhost:8000/ws"
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
	if runtime.GOOS != "linux" {
		log.Fatalf("Iskra Node only works on Linux machines currently")
	}

	e := echo.New()

	apiRoutes := e.Group("/api/v1")
	apiRoutes.GET("/meminfo", handlers.GetMemInfo) // GET memory information
	apiRoutes.GET("/cpuinfo", handlers.GetCpuInfo) // GET CPU information
	apiRoutes.GET("/uptime", handlers.GetUptime)   // GET system uptime

	// WebSocket endpoint for handling WebSocket connections
	e.GET("/ws", handlers.WebsocketHandler)

	go websocketClient()

	// Start the web server on port 8081.
	fmt.Println("Server started on :8081")
	if err := e.Start(":8081"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
