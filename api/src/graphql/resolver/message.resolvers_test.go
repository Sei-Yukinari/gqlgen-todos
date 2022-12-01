package resolver_test

import (
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func Test_mutationResolver_PostMessage(t *testing.T) {
	r := test.NewResolverMock(t, mysqlContainer, redisContainer)
	c := test.NewGqlgenClient(r)
	t.Run("Post Message", func(t *testing.T) {
		message := &model.Message{
			ID:        ksuid.New().String(),
			CreatedAt: time.Now().UTC(),
			User:      "Dummy User",
			Text:      "Dummy Text",
		}
		var res struct {
			PostMessage *gmodel.Message
		}
		// TODO time.Time not support https://github.com/99designs/gqlgen/issues/1372
		c.MustPost(`
			mutation($input: PostMessageInput) {
			  postMessage(input:$input) {
				id
				user
				text
			  }
}`,
			&res,
			client.Var("input", gmodel.PostMessageInput{
				User: "Dummy User",
				Text: "Dummy Text",
			}),
		)
		actual := res.PostMessage
		expected := r.Presenter.Message(message)
		assert.Equal(t, expected.Text, actual.Text)
		assert.Equal(t, expected.User, actual.User)
	})
}
