package test

import (
	"context"
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

func SetupRDB(t *testing.T, resource *dockertest.Resource) *gorm.DB {
	rdb := ConnectMySQLContainer(resource, pool)
	tx := rdb.Begin()
	t.Cleanup(func() {
		tx.Rollback()
	})
	return tx
}

func SetupRedis(t *testing.T, resource *dockertest.Resource) *redis.Client {
	cache := ConnectRedisContainer(resource, pool)
	t.Cleanup(func() {
		cache.FlushDB(context.Background())
	})
	return cache
}
