package server

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/resolver"
	"github.com/gin-gonic/gin"
)

func NewRouter(rsl *resolver.Resolver, middlewares []gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.GET("/", playgroundHandler())
	for _, m := range middlewares {
		r.Use(m)
	}
	r.POST("/query", graphqlHandler(rsl))
	r.GET("/query", graphqlHandler(rsl))

	return r
}
