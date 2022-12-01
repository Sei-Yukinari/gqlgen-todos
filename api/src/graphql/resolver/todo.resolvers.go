package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/Sei-Yukinari/gqlgen-todos/graph/generated"
	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	gerror "github.com/Sei-Yukinari/gqlgen-todos/src/graphql/error"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/loader"
	"github.com/Sei-Yukinari/gqlgen-todos/src/util/apperror"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input gmodel.NewTodo) (*gmodel.Todo, error) {
	userID, _ := strconv.Atoi(input.UserID)
	todo := model.Todo{
		Text:   input.Text,
		Done:   false,
		UserID: userID,
	}
	t, err := r.Repositories.Todo.Create(ctx, &todo)
	if err != nil {
		return nil, err
	}
	return r.Presenter.Todo(t), nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*gmodel.Todo, error) {
	todos, err := r.Repositories.Todo.FindAll(ctx)
	if err != nil {
		gerror.HandleError(ctx, apperror.Wrap(err))
		return nil, nil
	}
	return r.Presenter.Todos(todos), nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *gmodel.Todo) (*gmodel.User, error) {
	if obj.User == nil {
		return &gmodel.User{
			ID:   "0",
			Name: "",
		}, nil
	}
	user, err := loader.LoadUser(ctx, obj.User.ID)
	if err != nil {
		gerror.HandleError(ctx, apperror.Wrap(err))
		return nil, nil
	}
	if user == nil {
		return nil, nil
	}
	return &gmodel.User{
		ID:   obj.User.ID,
		Name: user.Name,
	}, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
