package gateway

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/repository"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
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

func (m Message) PostAndPublish(ctx context.Context, message *model.Message) (*model.Message, error) {
	messageJson, _ := json.Marshal(message)
	if err := m.redis.LPush(ctx, KeyMessages, string(messageJson)).Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	m.redis.Publish(ctx, PostMessagesSubscription, messageJson)
	return message, nil
}

func (m Message) FindAll(ctx context.Context) ([]*model.Message, error) {
	cmd := m.redis.LRange(ctx, KeyMessages, 0, -1)
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
	var messages []*model.Message
	for _, messageJson := range result {
		m := &model.Message{}
		_ = json.Unmarshal([]byte(messageJson), &m)
		messages = append(messages, m)
	}
	return messages, nil
}
