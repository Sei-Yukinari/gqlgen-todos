package usecase_test

import (
	"context"
	"testing"

	"github.com/Sei-Yukinari/gqlgen-todos/src/gateway"
	"github.com/Sei-Yukinari/gqlgen-todos/test"
	"github.com/ory/dockertest/v3"
)

var ctx context.Context
var mysqlContainer, redisContainer *dockertest.Resource
var repositories *gateway.Repositories

func TestMain(m *testing.M) {
	mysqlContainer = test.CreateMySQLContainer()
	redisContainer = test.CreateRedisContainer()
	m.Run()
	test.CloseContainer(mysqlContainer)
	test.CloseContainer(redisContainer)
}
