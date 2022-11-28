package gateway

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/repository"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
	"gorm.io/gorm"
)

type Repositories struct {
	Todo    repository.TodoRepository
	User    repository.User
	Message repository.MessageRepository
}

func NewRepositories(rdb *gorm.DB, redis *redis.Client) *Repositories {
	return &Repositories{
		Todo:    NewTodo(rdb),
		User:    NewUser(rdb),
		Message: NewMessage(redis),
	}
}
