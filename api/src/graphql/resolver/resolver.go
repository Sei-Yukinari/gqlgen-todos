package resolver

import (
	"sync"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/subscriber"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
	"github.com/Sei-Yukinari/gqlgen-todos/src/interfaces/presenter"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	rdb          *gorm.DB
	redisClient  *redis.Client
	subscribers  subscriber.Subscribers
	repositories *gateway.Repositories
	presenter    *presenter.Presenter
	messages     []*gmodel.Message
	mutex        sync.Mutex
}

func New(
	rdb *gorm.DB,
	redis *redis.Client,
	subscribers subscriber.Subscribers,
	repositories *gateway.Repositories,
	presenter *presenter.Presenter,
) *Resolver {
	return &Resolver{
		rdb:          rdb,
		redisClient:  redis,
		subscribers:  subscribers,
		repositories: repositories,
		presenter:    presenter,
		mutex:        sync.Mutex{},
	}
}
