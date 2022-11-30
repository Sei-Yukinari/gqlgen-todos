package server

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/loader"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/Sei-Yukinari/gqlgen-todos/src/middleware"
	"github.com/gin-gonic/gin"
)

func NewMiddleware(loaders *loader.Loaders) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.NewCors(),
		logger.InjectInContext(),
		loader.InjectInContext(loaders),
	}
}
