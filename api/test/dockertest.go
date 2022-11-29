package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

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

func CreateMySQLContainer(pool *dockertest.Pool, sqlFileNames []string) *dockertest.Resource {
	// mysql options
	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Platform:   "linux/x86_64",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + MysqlPassword,
			"MYSQL_DATABASE=" + MysqlDATABASE,
		},
		Mounts: mountsFile(sqlFileNames),
	}

	// start container
	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource
}

func mountsFile(files []string) []string {
	pwd, _ := os.Getwd()
	var m []string
	for _, v := range files {
		m = append(m,
			fmt.Sprintf(
				"%s/../../../mysql/init/%s:/docker-entrypoint-initdb.d/%s",
				pwd,
				v,
				v,
			),
		)
	}
	return m
}

func ConnectMySQLContainer(resource *dockertest.Resource, pool *dockertest.Pool, t *testing.T) *gorm.DB {

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
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return db
}

func CreateRedisContainer(pool *dockertest.Pool) *dockertest.Resource {
	resource, err := pool.Run("redis", "3.2", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource
}

func ConnectRedisContainer(resource *dockertest.Resource, pool *dockertest.Pool, t *testing.T) *redis.Client {
	var client *redis.Client
	if err := pool.Retry(func() error {
		client = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")),
		})

		return client.Ping(context.Background()).Err()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// When you're done, kill and remove the container
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	return nil
}