package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/segmentio/ksuid"
)

// PostMessage is the resolver for the postMessage field.
func (r *mutationResolver) PostMessage(ctx context.Context, input *gmodel.PostMessageInput) (*gmodel.Message, error) {
	message := &model.Message{
		ID:        ksuid.New().String(),
		CreatedAt: time.Now().UTC(),
		User:      input.User,
		Text:      input.Text,
	}

	r.repositories.Message.PostAndPublish(ctx, message)

	return r.presenter.Message(message), nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*gmodel.Message, error) {
	messages, err := r.repositories.Message.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return r.presenter.Messages(messages), nil
}

// MessagePosted is the resolver for the messagePosted field.
func (r *subscriptionResolver) MessagePosted(ctx context.Context, user string) (<-chan *gmodel.Message, error) {
	return r.subscribers.Message.Subscribe(ctx, user), nil
}
