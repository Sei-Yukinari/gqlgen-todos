package test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/resolver"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/subscriber"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/server"
	"github.com/Sei-Yukinari/gqlgen-todos/src/interfaces/presenter"
	"github.com/ory/dockertest/v3"
)

func NewResolverMock(t *testing.T, mysqlContainer, redisContainer *dockertest.Resource) *resolver.Resolver {
	rdb := NewRDB(t, mysqlContainer)
	redis := NewRedis(t, redisContainer)
	repositories := gateway.NewRepositories(rdb, redis)
	s := subscriber.New(repositories)
	p := presenter.New()
	return resolver.New(rdb, redis, s, repositories, p)
}

func NewGqlgenClient(r *resolver.Resolver) *client.Client {
	srv := server.NewGraphqlServer(r)
	return client.New(srv)
}
