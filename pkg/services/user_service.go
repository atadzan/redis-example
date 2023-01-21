package services

import (
	"context"
	"redis-example/models"
	"redis-example/pkg/repository"
)

type UserService struct {
	repos *repository.UserManager
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repos: repo.User,
	}
}

func (s *UserService) Add(ctx context.Context, key string, value models.User) (int64, error) {
	return s.repos.Add(ctx, key, value)
}

func (s *UserService) GetById(ctx context.Context, id string) (models.User, error) {
	return s.repos.GetById(ctx, id)
}