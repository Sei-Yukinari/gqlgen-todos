package repository

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/util/apperror"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, apperror.AppError)
	FindAll(ctx context.Context) ([]*model.Todo, apperror.AppError)
}
