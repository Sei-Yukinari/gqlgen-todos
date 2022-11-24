package server

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/resolver"
	"github.com/Sei-Yukinari/gqlgen-todos/src/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(resolver *resolver.Resolver) *gin.Engine {

	middlewares := []gin.HandlerFunc{
		middleware.NewCors(),
	}

	r := gin.Default()
	r.GET("/", playgroundHandler())
	for _, m := range middlewares {
		r.Use(m)
	}
	r.POST("/query", graphqlHandler(resolver))
	r.GET("/query", graphqlHandler(resolver))

	return r
}
