// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/resolver"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/subscriber"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/rdb"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/server"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitRouter() *gin.Engine {
	db := rdb.New()
	client := redis.New()
	subscribers := subscriber.New(client)
	repositories := gateway.NewRepositories(db)
	resolverResolver := resolver.New(db, client, subscribers, repositories)
	engine := server.NewRouter(resolverResolver)
	return engine
}

// wire.go:

var Set = wire.NewSet(infrastructure.Set, graphql.Set, gateway.Set)