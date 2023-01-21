package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func Ping(ctx context.Context, client *redis.Client) error {
	_, err := client.Ping(ctx).Result()
	return err
}

func NewRedisClient(ctx context.Context, host, password string, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})
	return rdb, Ping(ctx, rdb)
}
