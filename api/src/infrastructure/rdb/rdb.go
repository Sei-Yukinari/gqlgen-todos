package rdb

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Sei-Yukinari/gqlgen-todos/src/infrastructure/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var dbInstance *gorm.DB

func New() *gorm.DB {
	if dbInstance != nil {
		return dbInstance
	}
	conn, err := gorm.Open(mysql.Open(connString()), &gorm.Config{
		Logger: newLogger(),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		logger.Fatalf("db connection error: %v", err)
	} else {
		logger.Info("success to connect db!")
		dbConfig, _ := conn.DB()
		dbConfig.SetMaxOpenConns(0)
		dbConfig.SetMaxIdleConns(2)
		dbConfig.SetConnMaxLifetime(time.Hour * 1)
	}

	dbInstance = conn
	return dbInstance
}

func connString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"password",
		"mysql",
		"3306",
		"dev",
	)
}

func newLogger() gormlogger.Interface {
	conf := gormlogger.Config{
		SlowThreshold: time.Second,
		Colorful:      false,
		LogLevel:      gormlogger.Info,
	}

	return gormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		conf,
	)
}
