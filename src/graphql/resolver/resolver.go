package resolver

import gmodel "github.com/Sei-Yukinari/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*gmodel.Todo
}
