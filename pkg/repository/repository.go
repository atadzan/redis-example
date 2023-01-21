package repository

import "github.com/go-redis/redis/v8"

type Repository struct {
	User *UserManager
}

func NewRepository(client *redis.Client) *Repository {
	return &Repository{
		User: NewUserManager(client),
	}
}
