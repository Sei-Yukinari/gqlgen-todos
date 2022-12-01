package test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/Sei-Yukinari/gqlgen-todos/graph/generated"
	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/resolver"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/subscriber"
	"github.com/Sei-Yukinari/gqlgen-todos/src/interfaces/presenter"
	"github.com/ory/dockertest/v3"
)

func NewResolverMock(t *testing.T, mysqlContainer, redisContainer *dockertest.Resource) *resolver.Resolver {
	rdb := SetupRDB(t, mysqlContainer)
	redis := SetupRedis(t, redisContainer)
	repositories := gateway.NewRepositories(rdb, redis)
	s := subscriber.New(repositories)
	p := presenter.New()
	return resolver.New(rdb, redis, s, repositories, p)
}

func NewGqlgenClient(r *resolver.Resolver) *client.Client {
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: r},
		),
	)
	return client.New(srv)
}
