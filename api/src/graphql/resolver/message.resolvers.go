package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"log"
	"time"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
	"github.com/segmentio/ksuid"
)

// PostMessage is the resolver for the postMessage field.
func (r *mutationResolver) PostMessage(ctx context.Context, user string, text string) (*gmodel.Message, error) {
	message := &gmodel.Message{
		ID:        ksuid.New().String(),
		CreatedAt: time.Now().UTC(),
		User:      user,
		Text:      text,
	}

	messageJson, _ := json.Marshal(message)
	if err := r.redisClient.LPush(ctx, redis.KeyMessages, string(messageJson)).Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	r.redisClient.Publish(ctx, redis.PostMessagesSubscription, messageJson)

	return message, nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*gmodel.Message, error) {
	cmd := r.redisClient.LRange(ctx, redis.KeyMessages, 0, -1)
	err := cmd.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := cmd.Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var messages []*gmodel.Message
	for _, messageJson := range result {
		m := &gmodel.Message{}
		_ = json.Unmarshal([]byte(messageJson), &m)
		messages = append(messages, m)
	}

	return messages, nil
}

// MessagePosted is the resolver for the messagePosted field.
func (r *subscriptionResolver) MessagePosted(ctx context.Context, user string) (<-chan *gmodel.Message, error) {
	return r.subscribers.Message.Subscribe(ctx, user), nil
}
