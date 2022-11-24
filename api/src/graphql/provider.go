package graphql

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/resolver"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/subscriber"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	resolver.New,
	subscriber.New,
)
