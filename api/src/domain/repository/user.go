package repository

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/util/apperror"
)

type User interface {
	FindByIDs(ctx context.Context, ids []int) ([]*model.User, apperror.AppError)
}
