package resolver

import (
	"sync"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/subscriber"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos       []*gmodel.Todo
	redisClient *redis.Client
	subscribers subscriber.Subscribers
	messages    []*gmodel.Message
	mutex       sync.Mutex
}

func New(redis *redis.Client, subscribers subscriber.Subscribers) *Resolver {
	return &Resolver{
		redisClient: redis,
		subscribers: subscribers,
		mutex:       sync.Mutex{},
	}
}
