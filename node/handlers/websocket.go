package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// WebSocket upgrade handler
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Currently allowing all origins - this is not secure
		return true
	},
}

// Initalize the WebSocket connection and handle incoming messages
func WebsocketHandler(c echo.Context) error {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return err
	}
	defer conn.Close()

	fmt.Println("WebSocket client connected")

	// Handle incoming messages
	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Printf("Received from client: %s\n", message)

		// Echo the message back to the client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}

	return nil
}
