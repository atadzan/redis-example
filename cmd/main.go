package main

import (
	"context"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
	"redis-example/pkg/handlers"
	"redis-example/pkg/repository"
	"redis-example/pkg/services"
	"redis-example/redis"
)

func main() {
	ctx := context.Background()
	_ = initConfig()
	redisClient, err := redis.NewRedisClient(ctx, viper.GetString("redis.host")+":"+viper.GetString("redis.port"), "", viper.GetInt("redis.userDb"))
	if err != nil {
		log.Fatalln(err)
	}
	defer func(redisClient *redis2.Client) {
		err = redisClient.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(redisClient)
	repos := repository.NewRepository(redisClient)
	service := services.NewService(repos)
	handler := handlers.NewHandler(service)

	app := fiber.New()
	handler.InitRoutes(app)

	log.Fatalln(app.Listen(viper.GetString("http.host") + ":" + "8088"))
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("local")
	return viper.ReadInConfig()
}
