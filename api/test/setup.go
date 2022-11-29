package test

import (
	"log"
	"testing"

	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/redis"
	"github.com/ory/dockertest/v3"
	"gorm.io/gorm"
)

var pool *dockertest.Pool

func init() {
	pool = newPool()
}

func newPool() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return pool
}

func SetupRDB(t *testing.T) *gorm.DB {
	resource := CreateMySQLContainer(pool, []string{"todos.sql", "users.sql"})
	rdb := ConnectMySQLContainer(resource, pool, t)
	t.Cleanup(func() {
		pool.Purge(resource)
	})
	return rdb
}

func SetupRedis(t *testing.T) *redis.Client {
	resource := CreateRedisContainer(pool)
	redis := ConnectRedisContainer(resource, pool, t)
	t.Cleanup(func() {
		pool.Purge(resource)
	})
	return redis
}
