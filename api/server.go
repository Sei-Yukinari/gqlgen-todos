package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Sei-Yukinari/gqlgen-todos/graph/generated"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/resolver"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	redis := redis.New(context.Background())
	rl := resolver.NewResolver(redis)
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: rl,
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
