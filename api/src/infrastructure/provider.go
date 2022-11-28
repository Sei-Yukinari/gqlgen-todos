package infrastructure

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/rdb"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/server"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	redis.New,
	rdb.New,
	server.NewMiddleware,
	server.NewRouter,
)
