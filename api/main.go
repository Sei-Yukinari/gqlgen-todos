package main

import (
	"github.com/Sei-Yukinari/gqlgen-todos/di"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/server"
)

func main() {
	router := di.InitRouter()
	server.Run(router)
}
