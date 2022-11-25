package resolver

import (
	"sync"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/subscriber"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos        []*gmodel.Todo
	rdb          *gorm.DB
	redisClient  *redis.Client
	subscribers  subscriber.Subscribers
	repositories *gateway.Repositories
	messages     []*gmodel.Message
	mutex        sync.Mutex
}

func New(
	rdb *gorm.DB,
	redis *redis.Client,
	subscribers subscriber.Subscribers,
	repositories *gateway.Repositories,
) *Resolver {
	return &Resolver{
		rdb:          rdb,
		redisClient:  redis,
		subscribers:  subscribers,
		repositories: repositories,
		mutex:        sync.Mutex{},
	}
}
