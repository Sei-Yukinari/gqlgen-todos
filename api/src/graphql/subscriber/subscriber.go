package subscriber

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
)

type Subscribers struct {
	Message *MessageSubscriber
}

func New(repositories *gateway.Repositories) Subscribers {
	return Subscribers{
		Message: NewMessage(context.Background(), repositories.Message),
	}
}
