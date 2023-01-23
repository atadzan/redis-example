package repository

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"log"
	"redis-example/models"
	"time"
)

type UserManager struct {
	client *redis.Client
}

func NewUserManager(client *redis.Client) *UserManager {
	return &UserManager{
		client: client,
	}
}

func (r *UserManager) Add(ctx context.Context, score float64, key string, value models.User) (int64, error) {

	// Adding to sorted set
	resultCount, iErr := r.client.ZAdd(ctx, "users", []*redis.Z{{score, key}}...).Result()
	if iErr != nil {
		log.Fatalln(iErr)
	}
	sub, err := json.Marshal(value)
	if err != nil {
		log.Fatalln(err)
	}

	// Adding value to hash
	if err = r.client.Set(ctx, key, sub, 5*time.Minute).Err(); err != nil {

		log.Fatalln(err)
	}

	return resultCount, nil
}

func (r *UserManager) CountSetElems(ctx context.Context, key string) (int64, error) {

	// Get total count of elems in sorted set by set key
	total, tErr := r.client.ZCard(ctx, key).Result()
	if tErr != nil {
		log.Fatalln(tErr)
	}
	return total, nil
}

func (r *UserManager) GetAll(ctx context.Context, key string, offset, count int64) ([]models.User, error) {

	// Get list of elem keys from sorted set
	result, err := r.client.ZRangeByScore(ctx, key, &redis.ZRangeBy{Min: "-inf", Max: "+inf", Offset: offset, Count: count}).Result()
	if err != nil {
		log.Fatalln(err)
	}
	var users []models.User
	for _, k := range result {
		var user models.User
		// Get the value by key
		results, rErr := r.client.Get(ctx, k).Result()
		if rErr != nil {
			log.Fatalln(rErr)
		}
		err = jsoniter.UnmarshalFromString(results, user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserManager) GetByKey(ctx context.Context, key string) (models.User, error) {
	var user models.User

	// Get the value by key
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return user, err
	}
	err = jsoniter.UnmarshalFromString(result, &user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserManager) RemoveElemFromSet(ctx context.Context, set, key string) (int64, error) {
	// Remove key from sorted set
	result, err := r.client.ZRem(ctx, set, key).Result()
	if err != nil {
		log.Fatalln(err)
	}

	// Remove elem by key
	_ = r.client.GetDel(ctx, key)
	return result, nil
}
