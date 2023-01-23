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

func (s *UserService) Add(ctx context.Context, score float64, key string, value models.User) (int64, error) {
	return s.repos.Add(ctx, score, key, value)
}

func (s *UserService) CountSetElems(ctx context.Context, key string) (int64, error) {
	return s.repos.CountSetElems(ctx, key)
}

func (s *UserService) GetAll(ctx context.Context, key string, offset, count int64) ([]models.User, error) {
	return s.repos.GetAll(ctx, key, offset, count)
}

func (s *UserService) GetByKey(ctx context.Context, key string) (models.User, error) {
	return s.repos.GetByKey(ctx, key)
}

func (s *UserService) RemoveElemFromSet(ctx context.Context, set, key string) (int64, error) {
	return s.repos.RemoveElemFromSet(ctx, set, key)
}
