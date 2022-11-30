package redis

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/go-redis/redis/v8"
)

type Client = redis.Client
type PubSub = redis.PubSub

func New() *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	err := client.Ping(context.Background()).Err()

	if err != nil {
		logger.Fatalf("failed to connect redis:%v\n", err)
	}

	logger.Info("success to connect redis!")

	return client
}
