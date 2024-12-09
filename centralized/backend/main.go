package main

import (
	"iskra/centralized/internal/database"
	"iskra/centralized/internal/handlers"
	"iskra/centralized/internal/middlewares"

	"github.com/labstack/echo/v4"
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

func main() {
	e := echo.New()

	RegisterRoutes(e)

	database.Init()

	e.Logger.Fatal(e.Start(":8080"))
}

func RegisterRoutes(e *echo.Echo) {
	auth := e.Group("/auth")
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)

	protected := e.Group("/protected")
	protected.Use(middlewares.JWTMiddleware("secret"))
	protected.GET("", handlers.Restricted)
}
