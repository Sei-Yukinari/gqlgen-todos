package repository

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
)

type MessageRepository interface {
	PostAndPublish(ctx context.Context, message *model.Message) (*model.Message, error)
	FindAll(ctx context.Context) ([]*model.Message, error)
}
