package gateway

import (
	"context"
	"testing"

	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/ory/dockertest/v3"
)

var ctx context.Context
var mysqlContainer, redisContainer *dockertest.Resource

func TestMain(m *testing.M) {
	ctx = context.Background()
	mysqlContainer = test.CreateMySQLContainer([]string{"todos.sql", "users.sql"})
	redisContainer = test.CreateRedisContainer()
	m.Run()
	test.CloseContainer(mysqlContainer)
	test.CloseContainer(redisContainer)
}
