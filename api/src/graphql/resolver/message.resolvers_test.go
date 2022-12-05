package resolver_test

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
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

func Test_queryResolver_Messages(t *testing.T) {
	r := test.NewResolverMock(t, mysqlContainer, redisContainer)
	c := test.NewGqlgenClient(r)
	ctx := context.Background()
	messages := []*model.Message{
		{
			ID:        "1",
			CreatedAt: time.Now().UTC(),
			User:      "Dummy User1",
			Text:      "Dummy Text1",
		},
		{
			ID:        "2",
			CreatedAt: time.Now().UTC(),
			User:      "Dummy User2",
			Text:      "Dummy Text2",
		},
	}
	_, _ = r.Repositories.Message.PostAndPublish(ctx, messages[0])
	_, _ = r.Repositories.Message.PostAndPublish(ctx, messages[1])
	var res struct {
		Messages []*gmodel.Message
	}
	c.MustPost(`
			query findMessages {
			  messages {
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
	actual := res.Messages
	expected := r.Presenter.Messages(messages)
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].ID > expected[j].ID
	})
	assert.Equal(t, expected[0].ID, actual[0].ID)
	assert.Equal(t, expected[0].User, actual[0].User)
	assert.Equal(t, expected[0].Text, actual[0].Text)
	assert.Equal(t, len(expected), len(actual))
}

func Test_subscriptionResolver_MessagePosted(t *testing.T) {
	r := test.NewResolverMock(t, mysqlContainer, redisContainer)
	c := test.NewGqlgenClient(r)
	t.Run("Subscription Message", func(t *testing.T) {
		sub := c.Websocket(`
			subscription($user: String!) {
			  messagePosted(user: $user) {
				id
				user
				text
			  }
}`,
			client.Var("user", "Subscription User"),
		)
		defer func() {
			err := sub.Close()
			if err != nil {
				logger.Warn(err)
			}
		}()
		logger.Info("Subscribe!")
		var msg struct {
			resp struct {
				MessagePosted *gmodel.Message
			}
			err error
		}

		expected := gmodel.PostMessageInput{
			User: "Dummy User",
			Text: "Dummy Text",
		}

		var res struct {
			PostMessage *gmodel.Message
		}
		err := c.Post(`
				mutation($input: PostMessageInput) {
				  postMessage(input:$input) {
					id
					user
					text
				  }
}`,
			&res,
			client.Var("input", expected),
		)
		logger.Info("Publish!")
		assert.NoError(t, err)
		msg.err = sub.Next(&msg.resp)
		assert.NoError(t, msg.err, "sub.Next")
		assert.Equal(t, expected.User, msg.resp.MessagePosted.User)
		assert.Equal(t, expected.Text, msg.resp.MessagePosted.Text)
	})
}
