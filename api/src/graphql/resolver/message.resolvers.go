package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"time"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/segmentio/ksuid"
)

// PostMessage is the resolver for the postMessage field.
func (r *mutationResolver) PostMessage(ctx context.Context, user string, text string) (*gmodel.Message, error) {
	message := &gmodel.Message{
		ID:        ksuid.New().String(),
		CreatedAt: time.Now().UTC(),
		User:      user,
		Text:      text,
	}

	// 投稿されたメッセージを保存し、subscribeしている全てのコネクションにブロードキャスト
	r.mutex.Lock()
	r.messages = append(r.messages, message)
	for _, ch := range r.subscribers {
		ch <- message
	}
	r.mutex.Unlock()

	return message, nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*gmodel.Message, error) {
	return r.messages, nil
}

// MessagePosted is the resolver for the messagePosted field.
func (r *subscriptionResolver) MessagePosted(ctx context.Context, user string) (<-chan *gmodel.Message, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.subscribers[user]; ok {
		err := fmt.Errorf("`%s` has already been subscribed", user)
		log.Print(err.Error())
		return nil, err
	}

	// チャンネルを作成し、リストに登録
	ch := make(chan *gmodel.Message, 1)
	r.subscribers[user] = ch
	log.Printf("`%s` has been subscribed!", user)

	// コネクションが終了したら、このチャンネルを削除する
	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.subscribers, user)
		r.mutex.Unlock()
		log.Printf("`%s` has been unsubscribed.", user)
	}()

	return ch, nil
}
