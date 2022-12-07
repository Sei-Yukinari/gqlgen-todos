package integration_test

import (
	"net/http"
	"testing"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/steinfletcher/apitest"
)

func TestTodos(t *testing.T) {
	t.Run("Query Todos", func(t *testing.T) {
		apitest.New().
			Handler(newHandler(t)).
			Post("/query").
			GraphQLQuery(`
				query findTodos {
				  todos {
					id
					text
					done
					user{
					  id
					  name
					}
				  }
				}
`).
			Expect(t).
			Status(http.StatusOK).
			Body(`
					{
						"data": {
						"todos": []
						}	
					}`,
			).
			End()
	})
	t.Run("Mutation Create Todo", func(t *testing.T) {
		apitest.New().
			Handler(newHandler(t)).
			Post("/query").
			GraphQLRequest(apitest.GraphQLRequestBody{
				Query: `
					mutation ($input: NewTodo!) {
					  createTodo(input: $input) {
						text
						done
					  }
					}
`,
				Variables: map[string]interface{}{
					"input": gmodel.NewTodo{
						Text:   "todo",
						UserID: "1",
					},
				},
			}).
			Expect(t).
			Status(http.StatusOK).
			Body("{\"data\":{\"createTodo\":{\"text\":\"todo\",\"done\":false}}}").
			End()
	})
}
