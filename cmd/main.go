package main

import (
	"context"
	"fmt"
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
	fmt.Println(redisClient)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer redisClient.Close()
	repos := repository.NewRepository(redisClient)
	services := services.NewService(repos)
	handler := handlers.NewHandler(services)

	app := fiber.New()
	handler.InitRoutes(app)
	fmt.Println(viper.GetString("http.host") + ":" + viper.GetString("http.port"))
	log.Fatalln(app.Listen(viper.GetString("http.host") + ":" + "8088"))

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("local")
	return viper.ReadInConfig()
}
