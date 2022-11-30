package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/repository"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
)

type MessageSubscriber struct {
	repository repository.MessageRepository
	usersChan  map[string]chan<- *gmodel.Message
	mutex      sync.Mutex
}

func NewMessage(ctx context.Context, repository repository.MessageRepository) *MessageSubscriber {
	subscriber := &MessageSubscriber{
		repository: repository,
		usersChan:  map[string]chan<- *gmodel.Message{},
		mutex:      sync.Mutex{},
	}

	subscriber.Start(ctx)

	return subscriber
}

func (s *MessageSubscriber) Start(ctx context.Context) {
	logger := logger.FromContext(ctx)
	pubsub := s.repository.Subscribe(ctx)
	go func() {
		pubsubCh := pubsub.Channel()
		for msg := range pubsubCh {
			message := &gmodel.Message{}
			err := json.Unmarshal([]byte(msg.Payload), message)
			if err != nil {
				logger.Warn(err.Error())
				continue
			}
			s.mutex.Lock()
			for _, ch := range s.usersChan {
				ch <- message
			}
			s.mutex.Unlock()
		}
	}()
}

func (s *MessageSubscriber) Subscribe(ctx context.Context, user string) <-chan *gmodel.Message {
	logger := logger.FromContext(ctx)
	s.mutex.Lock()
	if _, ok := s.usersChan[user]; ok {
		err := fmt.Errorf("`%s` has already been subscribed", user)
		logger.Warnf(err.Error())
		return nil
	}
	s.mutex.Unlock()

	ch := make(chan *gmodel.Message, 1)
	s.usersChan[user] = ch
	logger.Debugf("`%s` has been subscribed!", user)

	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.usersChan, user)
		s.mutex.Unlock()
		logger.Infof("`%s` has been unsubscribed.", user)
	}()

	return ch
}
