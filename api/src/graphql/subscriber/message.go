package subscriber

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
)

type MessageSubscriber struct {
	client    *redis.Client
	usersChan map[string]chan<- *gmodel.Message
	mutex     sync.Mutex
}

func NewMessage(ctx context.Context, client *redis.Client) *MessageSubscriber {
	subscriber := &MessageSubscriber{
		client:    client,
		usersChan: map[string]chan<- *gmodel.Message{},
		mutex:     sync.Mutex{},
	}

	subscriber.Start(ctx)

	return subscriber
}

func (s *MessageSubscriber) Start(ctx context.Context) {
	pubsub := s.client.Subscribe(ctx, redis.PostMessagesSubscription)
	go func() {
		pubsubCh := pubsub.Channel()
		for msg := range pubsubCh {
			message := &gmodel.Message{}
			err := json.Unmarshal([]byte(msg.Payload), message)
			if err != nil {
				log.Printf(err.Error())
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
	s.mutex.Lock()
	if _, ok := s.usersChan[user]; ok {
		err := fmt.Errorf("`%s` has already been subscribed", user)
		log.Print(err.Error())
		return nil
	}
	s.mutex.Unlock()

	ch := make(chan *gmodel.Message, 1)
	s.usersChan[user] = ch
	log.Printf("`%s` has been subscribed!", user)

	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.usersChan, user)
		s.mutex.Unlock()
		log.Printf("`%s` has been unsubscribed.", user)
	}()

	return ch
}