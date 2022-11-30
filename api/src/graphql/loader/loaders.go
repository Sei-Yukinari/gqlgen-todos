package loader

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/dataloader"
)

type ctxKey string

const (
	loadersKey = ctxKey("data loaders")
)

// Loaders 各DataLoaderを取りまとめるstruct
type Loaders struct {
	UserLoader *dataloader.Loader
}

func New(repositories *gateway.Repositories) *Loaders {
	// define the data loader
	userLoader := &UserLoader{
		userRepository: repositories.User,
	}
	loaders := &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userLoader.BatchGetUsers),
	}
	return loaders
}

// InjectLoaders LoadersをcontextにインジェクトするHTTPミドルウェア
func InjectInContext(loaders *Loaders) gin.HandlerFunc {
	loaders.UserLoader.ClearAll()
	return func(c *gin.Context) {
		nextCtx := context.WithValue(c.Request.Context(), loadersKey, loaders)
		c.Request = c.Request.WithContext(nextCtx)
		c.Next()
	}
}

// GetLoaders ContextからLoadersを取得する
func FromContext(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
