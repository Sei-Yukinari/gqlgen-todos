package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Sei-Yukinari/gqlgen-todos/graph/generated"
	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
)

// Noop is the resolver for the noop field.
func (r *mutationResolver) Noop(ctx context.Context, input *gmodel.NoopInput) (*gmodel.NoopPayload, error) {
	panic(fmt.Errorf("not implemented: Noop - noop"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Noop is the resolver for the noop field.
func (r *subscriptionResolver) Noop(ctx context.Context, input *gmodel.NoopInput) (<-chan *gmodel.NoopPayload, error) {
	panic(fmt.Errorf("not implemented: Noop - noop"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
