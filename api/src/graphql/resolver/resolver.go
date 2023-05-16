package resolver

import (
	"sync"

	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/subscriber"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
	"github.com/Sei-Yukinari/gqlgen-todos/src/interfaces/presenter"
	"github.com/Sei-Yukinari/gqlgen-todos/src/usecase"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Rdb          *gorm.DB
	Redis        *redis.Client
	Subscribers  subscriber.Subscribers
	Repositories *gateway.Repositories
	UseCase      *usecase.UseCases
	Presenter    *presenter.Presenter
	mutex        sync.Mutex
}

func New(
	rdb *gorm.DB,
	redis *redis.Client,
	subscribers subscriber.Subscribers,
	repositories *gateway.Repositories,
	useCases *usecase.UseCases,
	presenter *presenter.Presenter,
) *Resolver {
	return &Resolver{
		Rdb:          rdb,
		Redis:        redis,
		Subscribers:  subscribers,
		Repositories: repositories,
		UseCase:      useCases,
		Presenter:    presenter,
		mutex:        sync.Mutex{},
	}
}
