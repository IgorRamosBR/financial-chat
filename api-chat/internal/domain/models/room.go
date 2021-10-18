package models

import (
	"api-chat/internal/application/dto/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name,omitempty"`
}

func ToRoom(room request.RoomRequest) Room {
	return Room{
		ID: primitive.NewObjectID(),
		Name: room.Name,
	}
}
