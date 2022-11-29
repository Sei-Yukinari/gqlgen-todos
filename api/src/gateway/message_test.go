package gateway

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/stretchr/testify/assert"
)

func TestMessage_PostAndPublish(t *testing.T) {
	t.Parallel()
	redis := test.SetupRedis(t)
	repo := NewMessage(redis)
	t.Run("Post And Publish", func(t *testing.T) {
		actual := &model.Message{
			ID:        "1",
			User:      "Dummy User",
			Text:      "Dummy",
			CreatedAt: time.Now(),
		}

		res, err := repo.PostAndPublish(ctx, actual)
		assert.NoError(t, err)
		assert.Equal(t, res, actual)
	})
}

func TestMessage_Subscribe(t *testing.T) {
	t.Parallel()
	redis := test.SetupRedis(t)
	repo := NewMessage(redis)

	t.Run("Subscribe", func(t *testing.T) {
		actual := &model.Message{
			ID:        "1",
			User:      "Dummy User",
			Text:      "Dummy",
			CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
		}

		pubsub := repo.Subscribe(ctx)

		_, err := repo.PostAndPublish(ctx, actual)
		assert.NoError(t, err)

		res := <-pubsub.Channel()
		expected := &model.Message{}
		err = json.Unmarshal([]byte(res.Payload), expected)
		if err != nil {
			log.Printf(err.Error())
		}
		assert.Equal(t, expected, actual)
	})
}

func TestMessage_FindAll(t *testing.T) {
	t.Parallel()
	redis := test.SetupRedis(t)
	repo := NewMessage(redis)
	t.Run("GET Message ALL", func(t *testing.T) {
		actual := []*model.Message{
			{
				ID:        "1",
				User:      "Dummy User",
				Text:      "Dummy",
				CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
			},
			{
				ID:        "2",
				User:      "Dummy User",
				Text:      "Dummy",
				CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, time.Local),
			},
		}

		_, err := repo.PostAndPublish(ctx, actual[0])
		assert.NoError(t, err)
		_, err = repo.PostAndPublish(ctx, actual[1])
		assert.NoError(t, err)

		msg, err := repo.FindAll(ctx)

		assert.ElementsMatch(t, msg, actual)
	})
}
