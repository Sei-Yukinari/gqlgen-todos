package resolver_test

import (
	"testing"

	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/ory/dockertest/v3"
)

var mysqlContainer, redisContainer *dockertest.Resource

func TestMain(m *testing.M) {
	mysqlContainer = test.CreateMySQLContainer()
	redisContainer = test.CreateRedisContainer()
	m.Run()
	test.CloseContainer(mysqlContainer)
	test.CloseContainer(redisContainer)
}
