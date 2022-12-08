package test

import (
	"context"
	"testing"

	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
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
		logger.Fatalf("Could not connect to docker: %s", err)
	}
	return pool
}

// NewRDB create mysql container
func NewRDB(t *testing.T, resource *dockertest.Resource) *gorm.DB {
	rdb := ConnectMySQLContainer(resource, pool)
	tx := rdb.Begin()
	t.Cleanup(func() {
		tx.Rollback()
	})
	return tx
}

// NewRedis create redis container
func NewRedis(t *testing.T, resource *dockertest.Resource) *redis.Client {
	cache := ConnectRedisContainer(resource, pool)
	t.Cleanup(func() {
		cache.FlushDB(context.Background())
	})
	return cache
}
