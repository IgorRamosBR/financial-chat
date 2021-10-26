package handlers

import (
	"api-chat/internal/application/dto/request"
	"api-chat/internal/domain/models"
	"api-chat/internal/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChatHandler struct {
	roomService services.RoomService
}

func NewChatHandler(roomService services.RoomService) ChatHandler {
	return ChatHandler{
		roomService: roomService,
	}
}

func (h ChatHandler) CreateRoom(c *gin.Context) {
	roomRequest, err := bindRoomRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room := models.ToRoom(roomRequest)
	err = h.roomService.CreateRoom(c, &room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func bindRoomRequest(c *gin.Context) (request.RoomRequest, error) {
	roomRequest := request.RoomRequest{}
	err := c.BindJSON(&roomRequest)
	if err != nil {
		return request.RoomRequest{}, err

	}
	return roomRequest, nil
}