package handlers

import (
	"api-chat/internal/application/dto/request"
	"api-chat/internal/domain/models"
	"api-chat/internal/domain/services"
	"github.com/labstack/echo/v4"
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

func (h ChatHandler) CreateRoom(ctx echo.Context) error {
	roomRequest, err := bindRoomRequest(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	room := models.ToRoom(roomRequest)
	err = h.roomService.CreateRoom(ctx.Request().Context(), &room)
	if err != nil {
		return err
	}

	return nil
}

func bindRoomRequest(c echo.Context) (request.RoomRequest, error) {
	roomRequest := request.RoomRequest{}
	err := c.Bind(&roomRequest)
	if err != nil {
		return request.RoomRequest{}, err

	}
	return roomRequest, nil
}