package presenter

import (
	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
)

func (p *Presenter) Message(message *model.Message) *gmodel.Message {
	return &gmodel.Message{
		ID:        message.ID,
		User:      message.User,
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
	}
}

func (p *Presenter) Messages(messages []*model.Message) []*gmodel.Message {
	var result []*gmodel.Message
	for _, v := range messages {
		result = append(result, p.Message(v))
	}
	return result
}
