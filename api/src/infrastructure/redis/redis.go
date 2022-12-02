package redis

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/config"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/go-redis/redis/v8"
)

type Client = redis.Client
type PubSub = redis.PubSub

func New() *Client {
	conf := config.Conf.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})
	err := client.Ping(context.Background()).Err()

	if err != nil {
		logger.Fatalf("failed to connect redis:%v\n", err)
	}

	logger.Info("success to connect redis!")

	return client
}
