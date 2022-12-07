package integration_test

import (
	"testing"

	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/loader"
	"github.com/Sei-Yukinari/gqlgen-todos/src/graphql/resolver"
	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/server"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/gin-gonic/gin"
	"github.com/ory/dockertest/v3"
)

var (
	mysqlContainer, redisContainer *dockertest.Resource
)

func TestMain(m *testing.M) {
	mysqlContainer = test.CreateMySQLContainer()
	redisContainer = test.CreateRedisContainer()
	m.Run()
	test.CloseContainer(mysqlContainer)
	test.CloseContainer(redisContainer)
}

func newHandler(t *testing.T) *gin.Engine {
	r := test.NewResolverMock(t, mysqlContainer, redisContainer)
	return initRouter(r)
}

func initRouter(r *resolver.Resolver) *gin.Engine {
	loaders := loader.New(r.Repositories)
	v := server.NewMiddleware(loaders)
	return server.NewRouter(r, v)
}
