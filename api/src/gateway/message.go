package gateway

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/repository"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
	"github.com/Sei-Yukinari/gqlgen-todos/src/util/apperror"
)

type Message struct {
	redis *redis.Client
}

func NewMessage(redis *redis.Client) *Message {
	return &Message{
		redis: redis,
	}
}

var _ repository.MessageRepository = (*Message)(nil)

const (
	PostMessagesSubscription = "messages"
	KeyMessages              = "messages-key"
)

func (m Message) PostAndPublish(ctx context.Context, message *model.Message) (*model.Message, apperror.AppError) {
	buf, _ := json.Marshal(message)
	if err := m.redis.LPush(ctx, KeyMessages, string(buf)).Err(); err != nil {
		return nil, apperror.Wrap(err)
	}
	m.publish(ctx, buf)
	return message, nil
}

func (m Message) publish(ctx context.Context, buf []byte) {
	m.redis.Publish(ctx, PostMessagesSubscription, buf)
}

func (m Message) Subscribe(ctx context.Context) *redis.PubSub {
	return m.redis.Subscribe(ctx, redis.PostMessagesSubscription)
}

func (m Message) FindAll(ctx context.Context) ([]*model.Message, apperror.AppError) {
	cmd := m.redis.LRange(ctx, KeyMessages, 0, -1)
	err := cmd.Err()
	if err != nil {
		return nil, apperror.Wrap(err)
	}

	result, err := cmd.Result()
	if err != nil {
		log.Println(err)
		return nil, apperror.Wrap(err)
	}
	var messages []*model.Message
	for _, messageJson := range result {
		m := &model.Message{}
		_ = json.Unmarshal([]byte(messageJson), &m)
		messages = append(messages, m)
	}
	return messages, nil
}
