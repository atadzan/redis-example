package repository

import (
	"context"
	"encoding/json"
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
	customJson, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}
	//data := map[string]interface{}{
	//	"name":     value.Name,
	//	"lastName": value.LastName,
	//	"age":      value.Age,
	//}
	//outputData, err := json.Marshal(data)
	//if err != nil {
	//	return 0, err
	//}
	fmt.Println(key)
	//result, iErr := r.client.Do(ctx, "ZADD", key, 2, customJson).Result()
	var members []*redis.Z
	member := redis.Z{Score: 0, Member: customJson}

	members = append(members, &member)
	result, iErr := r.client.ZAdd(ctx, key, members...).Result()
	if iErr != nil {
		return 0, err
	}
	fmt.Println(result)
	return 1, nil
}

func (r *UserManager) GetById(ctx context.Context, id string) (models.User, error) {
	var user models.User
	//result := r.client.Z
	return user, nil
}
