package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/graph/generated"
	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/loader"
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
	return r.presenter.Todo(t), nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*gmodel.Todo, error) {
	todos, err := r.repositories.Todo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return r.presenter.Todos(todos), nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *gmodel.Todo) (*gmodel.User, error) {
	user, err := loader.LoadUser(ctx, obj.User.ID)
	if err != nil {
		return nil, err
	}
	return &gmodel.User{
		ID:   obj.User.ID,
		Name: user.Name,
	}, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
