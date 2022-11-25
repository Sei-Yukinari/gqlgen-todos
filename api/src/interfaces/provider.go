package interfaces

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/interfaces/presenter"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	presenter.New,
)
