package server

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Sei-Yukinari/gqlgen-todos/graph/generated"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func NewGraphqlServer(resolver generated.ResolverRoot) *handler.Server {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	srv := handler.New(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: resolver,
		},
	))
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
	return srv
}

func graphqlHandler(resolver generated.ResolverRoot) gin.HandlerFunc {
	srv := NewGraphqlServer(resolver)
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
