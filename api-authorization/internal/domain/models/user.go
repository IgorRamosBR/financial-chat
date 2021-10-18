package models

import (
	"api-authorization/internal/application/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
}

func ToUser(user dto.UserRequest) User {
	return User{
		ID: primitive.NewObjectID(),
		Username: user.Username,
		Password: user.Password,
	}
}
