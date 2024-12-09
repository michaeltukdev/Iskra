package handlers

import (
	"context"
	"fmt"
	cpu "iskra/node/internal/cpuinfo"
	internal "iskra/node/internal/meminfo"
	uptime "iskra/node/internal/uptime"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

func Websocket() {
	fmt.Println("Websocket starting...")

	// Establish the WebSocket connection
	c, _, err := websocket.Dial(context.Background(), "ws://host.docker.internal:8081/hello", nil)
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer c.Close(websocket.StatusNormalClosure, "Normal closure")

	// Send messages
	messages := []interface{}{
		internal.Meminfo(),
		cpu.CPUInfo(),
		uptime.Uptime(),
	}

	for _, msg := range messages {
		// Create a fresh context for each message
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := wsjson.Write(ctx, c, msg)
		if err != nil {
			fmt.Printf("Failed to send message: %v\n", err)
			continue
		}
		fmt.Printf("Message sent: %v\n", msg)
	}
}

