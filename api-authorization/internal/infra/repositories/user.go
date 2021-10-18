package repositories

import (
	"api-authorization/internal/domain/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type UserRepositoryImpl struct {
	client *mongo.Client
}

func NewUserRepository(client *mongo.Client) UserRepository {
	return UserRepositoryImpl{client: client}
}

func (r UserRepositoryImpl) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.client.Database("authorization").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}