package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input gmodel.NewTodo) (*gmodel.Todo, error) {
	todo := model.Todo{
		Text: input.Text,
		Done: false,
	}
	t, err := r.repositories.Todo.Create(ctx, &todo)
	if err != nil {
		return nil, err
	}
	result := &gmodel.Todo{
		ID:   t.ID,
		Text: t.Text,
		Done: t.Done,
		User: &gmodel.User{
			ID:   "aaa",
			Name: "bbb",
		},
	}
	return result, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*gmodel.Todo, error) {
	todos, err := r.repositories.Todo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var result []*gmodel.Todo
	for _, v := range todos {
		t := &gmodel.Todo{
			ID:   v.ID,
			Text: v.Text,
			Done: false,
		}
		result = append(result, t)
	}
	return result, nil
}
