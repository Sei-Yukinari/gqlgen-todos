package resolver

import (
	"sync"

	gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos       []*gmodel.Todo
	subscribers map[string]chan<- *gmodel.Message
	messages    []*gmodel.Message
	mutex       sync.Mutex
}

func NewResolver() *Resolver {
	return &Resolver{
		subscribers: map[string]chan<- *gmodel.Message{},
		mutex:       sync.Mutex{},
	}
}
