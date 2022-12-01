package test

import (
	"context"

	"github.com/99designs/gqlgen/client"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/loader"
)

func InjectLoaderInContext(ctx context.Context, l *loader.Loaders) client.Option {
	return func(bd *client.Request) {
		ctx = context.WithValue(ctx, loader.LoadersKey, l)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	}
}
