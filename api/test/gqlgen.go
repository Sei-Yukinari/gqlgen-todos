package test

import (
	"net/http"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/Sei-Yukinari/gqlgen-todos/graph/generated"
	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/resolver"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/subscriber"
	"github.com/Sei-Yukinari/gqlgen-todos/src/interfaces/presenter"
	"github.com/gorilla/websocket"
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
	srv.Use(extension.Introspection{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			HandshakeTimeout: 10 * time.Second,
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	return client.New(srv)
}
