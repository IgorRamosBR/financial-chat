package services

import (
	"api-chat/internal/domain/models"
	"api-chat/internal/infra/repositories"
	"context"
)

type RoomService interface {
	CreateRoom(ctx context.Context, room *models.Room) error
}

type RoomServiceImpl struct {
	repository repositories.RoomRepository
}

func NewUserService(repository repositories.RoomRepository) RoomService {
	return RoomServiceImpl{
		repository: repository,
	}
}
func (s RoomServiceImpl) CreateRoom(ctx context.Context, room *models.Room) error {
	err := s.repository.SaveRoom(ctx, room)
	if err != nil {
		return err
	}
	return nil
}
