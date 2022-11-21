package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type Client = redis.Client

func New(ctx context.Context) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	err := client.Ping(ctx).Err()

	if err != nil {
		log.Fatalf("failed to connect redis:%v\n", err)
	}

	log.Println("success to connect redis!")

	return client
}
