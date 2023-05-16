package usecase

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/repository"
	gerror "github.com/Sei-Yukinari/gqlgen-todos/src/graphql/error"
	"github.com/Sei-Yukinari/gqlgen-todos/src/util/apperror"
)

type TodoUseCase interface {
	Get(context.Context) ([]*model.Todo, apperror.AppError)
}

type todoUsecase struct {
	repository repository.TodoRepository
}

func NewTodoUseCase(r repository.TodoRepository) TodoUseCase {
	return &todoUsecase{
		repository: r,
	}
}

func (tu todoUsecase) Get(ctx context.Context) ([]*model.Todo, apperror.AppError) {
	todos, err := tu.repository.FindAll(ctx)
	if err != nil {
		gerror.HandleError(ctx, apperror.Wrap(err))
		return nil, nil
	}
	return todos, nil
}
