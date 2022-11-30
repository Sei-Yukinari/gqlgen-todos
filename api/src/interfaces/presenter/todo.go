package presenter

import (
	"strconv"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
	"github.com/Sei-Yukinari/gqlgen-todos/src/domain/model"
)

func (p *Presenter) Todo(todo *model.Todo) *gmodel.Todo {
	var user *gmodel.User
	if todo.UserID == 0 {
		user = nil
	} else {
		user = &gmodel.User{
			ID:   strconv.Itoa(todo.UserID),
			Name: "",
		}
	}

	return &gmodel.Todo{
		ID:   strconv.Itoa(todo.ID),
		Text: todo.Text,
		Done: todo.Done,
		User: user,
	}
}

func (p *Presenter) Todos(todos []*model.Todo) []*gmodel.Todo {
	var result []*gmodel.Todo

	for _, v := range todos {
		result = append(result, p.Todo(v))
	}
	return result
}
