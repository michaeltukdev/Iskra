package main

import (
	"context"
	"log"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/labstack/echo/v4"
)

func handleWebSocket(c echo.Context) error {
	channel, err := websocket.Accept(c.Response(), c.Request(), nil)
	if err != nil {
		log.Printf("Failed to accept WebSocket: %v\n", err)
		return err
	}
	defer channel.Close(websocket.StatusNormalClosure, "Normal closure")

	for {
		var v interface{}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := wsjson.Read(ctx, channel, &v)
		if err != nil {
			switch websocket.CloseStatus(err) {
			case websocket.StatusNormalClosure, websocket.StatusGoingAway:
				return nil
			}

			log.Printf("Failed to read WebSocket message: %v\n", err)
			break
		}

		log.Printf("Received: %v\n", v)
	}

	return nil
}


func main() {
	e := echo.New()

	e.GET("/hello", handleWebSocket)

	log.Println("Starting server on :8081")
	e.Logger.Fatal(e.Start("0.0.0.0:8081"))
}
