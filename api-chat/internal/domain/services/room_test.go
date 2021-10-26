package services

import (
	"api-chat/internal/domain/models"
	"api-chat/internal/infra/repositories/mocks"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	room := &models.Room{
		Name: "Room 1",
	}
	roomRepository := mocks.NewMockRoomRepository(ctrl)
	roomService := NewRoomService(roomRepository)

	roomRepository.EXPECT().SaveRoom(gomock.Any(), gomock.Eq(room)).Return(nil)
	err := roomService.CreateRoom(context.TODO(), room)

	assert.Nil(t, err)
}

func TestCreateRoomError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	room := &models.Room{
		Name: "Room 1",
	}
	roomRepository := mocks.NewMockRoomRepository(ctrl)
	roomService := NewRoomService(roomRepository)

	roomRepository.EXPECT().SaveRoom(gomock.Any(), gomock.Eq(room)).Return(errors.New("internal error"))
	err := roomService.CreateRoom(context.TODO(), room)

	assert.Error(t, err, "internal error")
}