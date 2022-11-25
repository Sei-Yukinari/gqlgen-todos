package gateway

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	Todo repository.TodoRepository
}

func NewRepositories(rdb *gorm.DB) *Repositories {
	return &Repositories{
		Todo: NewTodo(rdb),
	}
}
