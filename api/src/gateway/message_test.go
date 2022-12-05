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
	r := test.SetupRedis(t, redisContainer)
	t.Run("Post And Publish", func(t *testing.T) {
		actual := &model.Message{
			ID:        "1",
			User:      "Dummy User",
			Text:      "Dummy",
			CreatedAt: time.Now(),
		}
		repo := gateway.NewMessage(r)
		res, err := repo.PostAndPublish(ctx, actual)
		assert.NoError(t, err)
		assert.Equal(t, res, actual)
	})
}

func TestMessage_Subscribe(t *testing.T) {
	r := test.SetupRedis(t, redisContainer)
	actual := &model.Message{
		ID:        "1",
		User:      "Dummy User",
		Text:      "Dummy",
		CreatedAt: time.Now().UTC(),
	}
	repo := gateway.NewMessage(r)
	pubsub := repo.Subscribe(ctx)
	defer func() {
		_ = pubsub.Close()
	}()
	logger.Info("Subscribe!")
	_, apperr := repo.PostAndPublish(ctx, actual)
	logger.Info("Publish!")
	assert.NoError(t, apperr)
	t.Run("Subscribe", func(t *testing.T) {
		select {
		case res := <-pubsub.Channel():
			logger.Infof("Received message!%+v", res)
			expected := &model.Message{}
			err := json.Unmarshal([]byte(res.Payload), expected)
			if err != nil {
				logger.Warn(err.Error())
			}
			assert.Equal(t, expected, actual)
		}
	})
}

func TestMessage_FindAll(t *testing.T) {
	r := test.SetupRedis(t, redisContainer)
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
		repo := gateway.NewMessage(r)
		_, err := repo.PostAndPublish(ctx, actual[0])
		assert.NoError(t, err)
		_, err = repo.PostAndPublish(ctx, actual[1])
		assert.NoError(t, err)

		msg, err := repo.FindAll(ctx)

		assert.Equal(t, len(msg), len(actual))
		assert.ElementsMatch(t, msg, actual)
	})
}
