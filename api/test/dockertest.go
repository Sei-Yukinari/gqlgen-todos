package test

import (
	"context"
	"fmt"
	"time"

	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"github.com/Sei-Yukinari/gqlgen-todos/src/path"
	"github.com/go-redis/redis/v8"
	"github.com/ory/dockertest/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MysqlUser     string = "root"
	MysqlPassword string = "password"
	MysqlDATABASE string = "test"
)

func CreateMySQLContainer() *dockertest.Resource {
	// mysql options
	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Platform:   "linux/x86_64",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + MysqlPassword,
			"MYSQL_DATABASE=" + MysqlDATABASE,
		},
		Mounts: []string{
			fmt.Sprintf(
				"%s:/docker-entrypoint-initdb.d/",
				path.GetProjectRootPath()+"/test/init/",
			),
		},
	}

	// start container
	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		logger.Fatalf("Could not start resource: %s", err)
	}

	return resource
}

func ConnectMySQLContainer(resource *dockertest.Resource, pool *dockertest.Pool) *gorm.DB {
	var db *gorm.DB
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@(localhost:%s)/%s?parseTime=true",
		MysqlUser,
		MysqlPassword,
		resource.GetPort("3306/tcp"),
		MysqlDATABASE,
	)
	if err := pool.Retry(func() error {
		// wait for container setup
		time.Sleep(time.Second * 3)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Ping()
	}); err != nil {
		logger.Fatalf("Could not connect to docker: %s", err)
	}
	return db
}

func CreateRedisContainer() *dockertest.Resource {
	resource, err := pool.Run("redis", "6.2", nil)
	if err != nil {
		logger.Fatalf("Could not start resource: %s", err)
	}
	return resource
}

func ConnectRedisContainer(resource *dockertest.Resource, pool *dockertest.Pool) *redis.Client {
	var client *redis.Client
	if err := pool.Retry(func() error {
		client = redis.NewClient(
			&redis.Options{
				Addr: fmt.Sprintf("localhost:%s",
					resource.GetPort("6379/tcp")),
			})

		return client.Ping(context.Background()).Err()
	}); err != nil {
		logger.Fatalf("Could not connect to docker: %s", err)
	}
	return client
}

func CloseContainer(resource *dockertest.Resource) {
	// stop container
	if err := pool.Purge(resource); err != nil {
		logger.Fatalf("Could not purge resource: %s", err)
	}
}
