package gateway

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/repository"
	"github.com/Sei-Yukinari/gqlgen-todos/src/util/apperror"
	"gorm.io/gorm"
)

type Todo struct {
	tx *gorm.DB
}

func NewTodo(tx *gorm.DB) *Todo {
	return &Todo{
		tx: tx,
	}
}

var _ repository.TodoRepository = (*Todo)(nil)

func (t Todo) Create(ctx context.Context, todo *model.Todo) (*model.Todo, apperror.AppError) {
	if err := t.tx.Create(todo).Error; err != nil {
		return nil, apperror.Wrap(err).SetCode(apperror.Database)
	}
	return todo, nil
}

func (t Todo) FindAll(ctx context.Context) ([]*model.Todo, apperror.AppError) {
	var todos []*model.Todo
	if err := t.tx.
		Find(&todos).
		Error; err != nil {
		return nil, apperror.Wrap(err).SetCode(apperror.Database)
	}
	return todos, nil
}
