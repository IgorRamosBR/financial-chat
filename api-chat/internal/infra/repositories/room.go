package repositories

import (
	"api-chat/internal/domain/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomRepository interface {
	SaveRoom(ctx context.Context, user *models.Room) error
}

type RoomRepositoryImpl struct {
	client *mongo.Client
}

func NewRoomRepository(client *mongo.Client) RoomRepository {
	return RoomRepositoryImpl{client: client}
}

func (r RoomRepositoryImpl) SaveRoom(ctx context.Context, user *models.Room) error {
	_, err := r.client.Database("rooms").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}