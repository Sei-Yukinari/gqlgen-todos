package subscriber

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
)

type Subscribers struct {
	Message *MessageSubscriber
}

func New(redis *redis.Client) Subscribers {
	return Subscribers{
		Message: NewMessage(context.Background(), redis),
	}
}
