package rdb

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Sei-Yukinari/gqlgen-todos/src/config"
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
	conf := config.Conf.Db
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
}

func newLogger() gormlogger.Interface {
	conf := gormlogger.Config{
		SlowThreshold: time.Second,
		Colorful:      true,
		LogLevel:      gormlogger.Info,
	}

	return gormlogger.New(
		log.New(os.Stdout, "[GORM]", log.Lmsgprefix),
		conf,
	)
}
