package gateway

import (
	"testing"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/stretchr/testify/assert"
)

func TestTodo_Create(t *testing.T) {
	rdb := test.SetupRDB(t, mysqlContainer)
	t.Run("Create TODO", func(t *testing.T) {
		actual := &model.Todo{
			ID:     1,
			Text:   "Dummy",
			Done:   true,
			UserID: 1,
		}
		repo := NewTodo(rdb)
		res, err := repo.Create(ctx, actual)
		assert.NoError(t, err)
		assert.Equal(t, res, actual)
	})

}

func TestTodo_FindAll(t *testing.T) {
	rdb := test.SetupRDB(t, mysqlContainer)
	t.Run("Get TODO ALL", func(t *testing.T) {
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
		repo := NewTodo(rdb)
		res, err := repo.FindAll(ctx)
		assert.Equal(t, err, nil)
		assert.Equal(t, len(res), len(actual))
	})
}
