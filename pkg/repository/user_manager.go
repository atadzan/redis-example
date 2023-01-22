package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
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

func (r *UserManager) Add(ctx context.Context, key string, value models.User) (int64, error) {
	var members []*redis.Z
	member := redis.Z{Score: 1, Member: key}
	members = append(members, &member)
	resultCount, iErr := r.client.ZAdd(ctx, "videos", members...).Result()
	if iErr != nil {
		fmt.Println(iErr.Error())
		return 0, iErr
	}

	// Another way of adding to cache
	//result, iErr := r.client.Do(ctx, "ZADD", key, 2, customJson).Result()

	if resultCount != 0 {
		//if _, err := r.client.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		//	rdb.
		//}); err != nil {
		//	return 0, err
		//}
		err := r.client.HSet(ctx, key, "name", value.Name, "lastName", value.LastName, "age", value.Age).Err()
		if err != nil {
			return 0, err
		}
		//fmt.Println()
	}

	return resultCount, nil
}

func (r *UserManager) Get(ctx context.Context, offset int64) ([]models.User, error) {
	fmt.Println("get repo")
	var users []models.User
	count := 5
	result, err := r.client.ZRangeByScore(ctx, "videos", &redis.ZRangeBy{Min: "0", Max: "-1", Offset: offset, Count: int64(count)}).Result()
	if err != nil {
		fmt.Println("Error 1")
		return users, err
	}
	fmt.Println("result: ", result)
	for _, key := range result {
		var user models.User
		err = r.client.HGetAll(ctx, key).Scan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
