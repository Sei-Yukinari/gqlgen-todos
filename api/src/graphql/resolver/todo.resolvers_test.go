package resolver_test

import (
	"testing"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/stretchr/testify/assert"
)

func Test_queryResolver_Todos(t *testing.T) {
	c, r := test.GqlgenClient(t, mysqlContainer, redisContainer)
	t.Run("Query Todos", func(t *testing.T) {
		q := `
		query findTodos {
		  todos {
			id
			text
			done
		  }
		}
`
		todos := []*model.Todo{
			{
				ID:   1,
				Text: "Dummy",
				Done: true,
			},
			{
				ID:   2,
				Text: "Dummy",
				Done: false,
			},
		}

		err := test.Seeds(r.Rdb,
			[]interface{}{
				todos,
			})
		if err != nil {
			logger.Fatalf("fail seed data: %s", err)
		}
		var res struct {
			Todos []*gmodel.Todo
		}
		c.MustPost(q, &res)
		actual := res.Todos
		expected := r.Presenter.Todos(todos)
		assert.Equal(t, expected, actual)
	})
}
