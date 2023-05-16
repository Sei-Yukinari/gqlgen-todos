package usecase_test

import (
	"testing"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/Sei-Yukinari/gqlgen-todos/src/usecase"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/stretchr/testify/assert"
)

func TestTodo_Get(t *testing.T) {
	rdb := test.NewRDB(t, mysqlContainer)
	redis := test.NewRedis(t, redisContainer)
	repositories := gateway.NewRepositories(rdb, redis)
	t.Run("Get", func(t *testing.T) {
		actual := []model.Todo{
			{
				ID:     1,
				Text:   "Dummy",
				Done:   true,
				UserID: 1,
			},
			{
				ID:     2,
				Text:   "Dummy",
				Done:   false,
				UserID: 1,
			},
		}

		err := test.Seeds(rdb,
			[]interface{}{
				actual,
			})
		if err != nil {
			logger.Fatalf("fail seed data: %s", err)
		}
		u := usecase.NewTodoUseCase(repositories.Todo)
		res, err := u.Get(ctx)
		assert.Equal(t, err, nil)
		assert.Equal(t, len(res), len(actual))
	})
}
