package resolver_test

import (
	"context"
	"testing"

	"github.com/99designs/gqlgen/client"
	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/loader"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/stretchr/testify/assert"
)

func Test_queryResolver_Todos(t *testing.T) {
	r := test.NewResolverMock(t, mysqlContainer, redisContainer)
	c := test.NewGqlgenClient(r)
	t.Run("Query Todos", func(t *testing.T) {
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
		c.MustPost(`
			query findTodos {
			  todos {
				id
				text
				done
			  }
			}`,
			&res,
		)
		actual := res.Todos
		expected := r.Presenter.Todos(todos)
		assert.Equal(t, expected, actual)
	})
}

func Test_mutationResolver_CreateTodo(t *testing.T) {
	r := test.NewResolverMock(t, mysqlContainer, redisContainer)
	c := test.NewGqlgenClient(r)
	ctx := context.Background()
	loaders := loader.New(r.Repositories)
	t.Run("Query Todos", func(t *testing.T) {
		todo := &model.Todo{
			ID:     1,
			Text:   "todo",
			Done:   false,
			UserID: 1,
		}
		user := []*model.User{
			{
				ID:   1,
				Name: "", //TODO set name
			},
		}
		err := test.Seeds(r.Rdb,
			[]interface{}{
				user,
			})
		if err != nil {
			logger.Fatalf("fail seed data: %s", err)
		}
		var res struct {
			CreateTodo *gmodel.Todo
		}
		c.MustPost(`
			mutation ($input: NewTodo!) {
			  createTodo(input: $input) {
				id
				text
				done
				user {
				  id
				  name
				}
			  }
			}`,
			&res,
			client.Var("input", gmodel.NewTodo{
				Text:   "todo",
				UserID: "1",
			}),
			test.InjectLoaderInContext(ctx, loaders),
		)
		actual := res.CreateTodo
		expected := r.Presenter.Todo(todo)
		assert.Equal(t, expected.Text, actual.Text)
		assert.Equal(t, expected.Done, actual.Done)
		assert.Equal(t, expected.User, actual.User)
	})
}
