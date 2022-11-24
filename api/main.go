package main

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/resolver"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/server"
)

func main() {
	redis := redis.New(context.Background())
	resolver := resolver.NewResolver(redis)

	router := server.NewRouter(resolver)
	server.Run(router)
}
