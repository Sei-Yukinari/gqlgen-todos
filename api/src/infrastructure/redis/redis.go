package redis

import (
	"context"
	"log"

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
		log.Fatalf("failed to connect redis:%v\n", err)
	}

	log.Println("success to connect redis!")

	return client
}
