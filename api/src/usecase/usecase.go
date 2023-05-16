package usecase

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
)

type UseCases struct {
	Todo TodoUseCase
}

func NewUseCases(repositories *gateway.Repositories) *UseCases {
	return &UseCases{
		Todo: NewTodoUseCase(repositories.Todo),
	}
}
