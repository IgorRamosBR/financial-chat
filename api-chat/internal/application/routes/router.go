package routes

import (
	"api-chat/internal/application/handlers"
	"api-chat/internal/domain/services"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{}
)

func NewRouter(roomService services.RoomService) *echo.Echo {
	e := echo.New()

	chatHandler := handlers.NewChatHandler(roomService)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/rooms", chatHandler.CreateRoom)
	e.File("/", "public/index.html")
	e.GET("/ws", hello)
	e.GET("/rooms/:id/messages", hello)

	return e
}

func hello(c echo.Context) error {
	id := c.Param("id")
	c.Logger().Info(id)
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	_, msg, err := ws.ReadMessage()
	if err != nil {
		c.Logger().Error(err)
	}
	fmt.Printf("%s\n", msg)
/*	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}*/
	return nil
}