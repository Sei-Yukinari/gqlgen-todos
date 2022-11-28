package repository

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
)

type User interface {
	FindByIDs(ctx context.Context, ids []int) ([]*model.User, error)
}
