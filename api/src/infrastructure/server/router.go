package server

import (
	"github.com/Sei-Yukinari/gqlgen-todos/graph/generated"
	"github.com/Sei-Yukinari/gqlgen-todos/src/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(resolver generated.ResolverRoot) *gin.Engine {

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
