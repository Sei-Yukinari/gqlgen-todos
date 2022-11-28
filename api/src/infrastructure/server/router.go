package server

import (
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/loader"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/resolver"
	"github.com/Sei-Yukinari/gqlgen-todos/src/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(rsl *resolver.Resolver) *gin.Engine {

	middlewares := []gin.HandlerFunc{
		middleware.NewCors(),
		//TODO
		loader.InjectLoaders(loader.NewLoaders(rsl.Rdb)),
	}

	r := gin.Default()
	r.GET("/", playgroundHandler())
	for _, m := range middlewares {
		r.Use(m)
	}
	r.POST("/query", graphqlHandler(rsl))
	r.GET("/query", graphqlHandler(rsl))

	return r
}
