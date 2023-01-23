package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"redis-example/models"
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

	if resultCount != 0 {
		// Adding value to hash
		if err := r.client.HSet(ctx, key, "name", value.Name, "lastName", value.LastName, "age", value.Age).Err(); err != nil {
			log.Fatalln(err)
		}
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

		// Get the value from hash by key
		err = r.client.HGetAll(ctx, k).Scan(&user)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserManager) GetByKey(ctx context.Context, key string) (models.User, error) {
	var user models.User

	// Get the value of key from hash & scan it into struct
	if err := r.client.HGetAll(ctx, key).Scan(&user); err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserManager) RemoveElemFromSet(ctx context.Context, set, key string) (int64, error) {
	// Remove key from sorted set
	result, err := r.client.ZRem(ctx, set, key).Result()
	if err != nil {
		log.Fatalln(err)
	}

	// Get all fields of elem by hash key
	keys, keyErr := r.client.HKeys(ctx, key).Result()
	if keyErr != nil {
		return 0, keyErr
	}

	// Remove elems all fields by hash key
	if err = r.client.HDel(ctx, key, keys...).Err(); err != nil {
		return 0, err
	}
	return result, nil
}
