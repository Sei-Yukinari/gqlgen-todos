package test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MysqlUser     string = "root"
	MysqlPassword string = "password"
	MysqlDATABASE string = "test"
)

func CreateMySQLContainer(sqlFileName string) (*dockertest.Resource, *dockertest.Pool) {
	// connect to docker
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	pwd, _ := os.Getwd()
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
			//migration
			pwd + "/" + sqlFileName + ":/docker-entrypoint-initdb.d/todos.sql",
		},
	}

	// start container
	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource, pool
}

func closeMySQLContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	// stop container
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
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
		time.Sleep(time.Second * 5)
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
	t.Cleanup(func() {
		closeMySQLContainer(resource, pool)
	})
	return db
}
