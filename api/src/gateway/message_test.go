package gateway_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/stretchr/testify/assert"
)

func TestMessage_PostAndPublish(t *testing.T) {
	redis := test.SetupRedis(t, redisContainer)
	t.Run("Post And Publish", func(t *testing.T) {
		actual := &model.Message{
			ID:        "1",
			User:      "Dummy User",
			Text:      "Dummy",
			CreatedAt: time.Now(),
		}
		repo := gateway.NewMessage(redis)
		res, err := repo.PostAndPublish(ctx, actual)
		assert.NoError(t, err)
		assert.Equal(t, res, actual)
	})
}

func TestMessage_Subscribe(t *testing.T) {
	redis := test.SetupRedis(t, redisContainer)
	t.Run("Subscribe", func(t *testing.T) {
		actual := &model.Message{
			ID:        "1",
			User:      "Dummy User",
			Text:      "Dummy",
			CreatedAt: time.Now().UTC(),
		}
		repo := gateway.NewMessage(redis)
		pubsub := repo.Subscribe(ctx)

		_, apperr := repo.PostAndPublish(ctx, actual)
		assert.NoError(t, apperr)

		res := <-pubsub.Channel()
		expected := &model.Message{}
		err := json.Unmarshal([]byte(res.Payload), expected)
		if err != nil {
			logger.Warn(err.Error())
		}
		assert.Equal(t, expected, actual)
	})
}

func TestMessage_FindAll(t *testing.T) {
	redis := test.SetupRedis(t, redisContainer)
	t.Run("GET Message ALL", func(t *testing.T) {
		actual := []*model.Message{
			{
				ID:        "1",
				User:      "Dummy User",
				Text:      "Dummy",
				CreatedAt: time.Now().UTC(),
			},
			{
				ID:        "2",
				User:      "Dummy User",
				Text:      "Dummy",
				CreatedAt: time.Now().UTC(),
			},
		}
		repo := gateway.NewMessage(redis)
		_, err := repo.PostAndPublish(ctx, actual[0])
		assert.NoError(t, err)
		_, err = repo.PostAndPublish(ctx, actual[1])
		assert.NoError(t, err)

		msg, err := repo.FindAll(ctx)

		assert.Equal(t, len(msg), len(actual))
		assert.ElementsMatch(t, msg, actual)
	})
}
