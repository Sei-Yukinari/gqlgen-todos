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

func GqlgenClient(t *testing.T, mysqlContainer, redisContainer *dockertest.Resource) (*client.Client, *resolver.Resolver) {
	rdb := SetupRDB(t, mysqlContainer)
	redis := SetupRedis(t, redisContainer)
	repositories := gateway.NewRepositories(rdb, redis)
	s := subscriber.New(repositories)
	p := presenter.New()
	r := resolver.New(rdb, redis, s, repositories, p)

	c := client.New(
		handler.NewDefaultServer(
			generated.NewExecutableSchema(generated.Config{Resolvers: r}),
		),
	)
	return c, r
}
