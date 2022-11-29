package gateway

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/stretchr/testify/assert"
)

var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	code := m.Run()

	os.Exit(code)
}

func TestTodo_Create(t *testing.T) {
	rdb := test.SetupRDB(t)
	t.Run("Create TODO", func(t *testing.T) {
		t.Parallel()
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
	t.Parallel()
	rdb := test.SetupRDB(t)
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

		test.Seeds(rdb,
			[]interface{}{
				actual,
			})
		repo := NewTodo(rdb)
		res, err := repo.FindAll(ctx)
		assert.Equal(t, err, nil)
		fmt.Println(len(res))
		fmt.Println(len(actual))

		assert.Equal(t, len(res), len(actual))
	})
}
