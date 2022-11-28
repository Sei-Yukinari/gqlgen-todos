package gateway

import (
	"context"
	"fmt"
	"testing"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setup(t *testing.T) (context.Context, *gorm.DB, *Todo) {
	resource, pool := test.CreateMySQLContainer("../../../mysql/init/todos.sql")
	ctx := context.Background()
	rdb := test.ConnectMySQLContainer(resource, pool, t)
	repo := NewTodo(rdb)
	return ctx, rdb, repo
}

func TestTodo_Create(t *testing.T) {
	t.Parallel()
	ctx, _, repo := setup(t)
	t.Run("Create User", func(t *testing.T) {
		actual := &model.Todo{
			ID:     1,
			Text:   "Dummy",
			Done:   true,
			UserID: 1,
		}
		res, err := repo.Create(ctx, actual)
		assert.NoError(t, err)
		assert.Equal(t, res, actual)
	})
}

func TestTodo_FindAll(t *testing.T) {
	t.Parallel()
	ctx, rdb, repo := setup(t)
	t.Run("Get User", func(t *testing.T) {
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

		test.Seeds(rdb,
			[]interface{}{
				actual,
			})

		res, err := repo.FindAll(ctx)
		assert.Equal(t, err, nil)
		fmt.Println(len(res))
		fmt.Println(len(actual))

		assert.Equal(t, len(res), len(actual))
	})
}
