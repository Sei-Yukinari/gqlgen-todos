package loader

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/util/apperror"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/dataloader"
)

type ctxKey string

const (
	LoadersKey = ctxKey("data loaders")
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
		nextCtx := context.WithValue(c.Request.Context(), LoadersKey, loaders)
		c.Request = c.Request.WithContext(nextCtx)
		c.Next()
	}
}

// GetLoaders ContextからLoadersを取得する
func FromContext(ctx context.Context) (*Loaders, apperror.AppError) {
	v := ctx.Value(LoadersKey)
	l, ok := v.(*Loaders)
	if !ok {
		return nil, apperror.New("failed to get loader from current context")
	}
	return l, nil
}
