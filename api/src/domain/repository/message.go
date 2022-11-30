package repository

import (
	"context"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
	"github.com/Sei-Yukinari/gqlgen-todos/src/util/apperror"
)

type MessageRepository interface {
	PostAndPublish(ctx context.Context, message *model.Message) (*model.Message, apperror.AppError)
	Subscribe(ctx context.Context) *redis.PubSub
	FindAll(ctx context.Context) ([]*model.Message, apperror.AppError)
}
