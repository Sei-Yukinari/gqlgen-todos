//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure"
	"github.com/Sei-Yukinari/gqlgen-todos/src/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	infrastructure.Set,
	interfaces.Set,
	graphql.Set,
	gateway.Set,
)

func InitRouter() *gin.Engine {
	wire.Build(Set)
	return &gin.Engine{}
}
