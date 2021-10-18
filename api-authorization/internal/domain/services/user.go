package services

import (
	"api-authorization/internal/domain/models"
	"api-authorization/internal/infra/repositories"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type UserServiceImpl struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return UserServiceImpl{
		repository: repository,
	}
}
func (s UserServiceImpl) CreateUser(ctx context.Context, user *models.User) error {
	err := s.repository.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
