package routes

import (
	"api-chat/internal/application/handlers"
	"api-chat/internal/domain/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{}
)

func NewRouter(roomService services.RoomService) *gin.Engine {
	lobby := services.NewLobby()
	go lobby.Run()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	chatHandler := handlers.NewChatHandler(roomService)

	r.POST("/rooms", chatHandler.CreateRoom)
	r.LoadHTMLFiles("public/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/ws", func(c *gin.Context) {
		serverWs(lobby, c.Writer, c.Request)
	})
	/*r.Static("/", "public/index.html")
	e.GET("/ws", hello)
	e.GET("/rooms/:id/messages", hello)*/

	return r
}

func serverWs(lobby *services.Lobby, w http.ResponseWriter, r *http.Request) {
	//id := c.Param("id")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf(err.Error())
	}

	user := services.NewUser(lobby, ws, make(chan []byte, 256))
	user.Lobby.Register <- user

	go user.WriteMessages()
	go user.ReadMessages()
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
}