package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App   App   `toml:"app"`
	Db    DB    `toml:"db"`
	Redis Redis `toml:"redis"`
}

type App struct {
	Port                        string `toml:"port"`
	TimeoutToGracefulShutdownMs int    `toml:"timeoutToGracefulShutdownMs"`
	CorsOrigins                 string `toml:"corsOrigins"`
}

type DB struct {
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"name"`
	Port     string `toml:"port"`
}

type Redis struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

var Conf *Config

func init() {
	Conf = new(Config)

	GoEnv := os.Getenv("GO_ENV")
	if GoEnv == "" {
		GoEnv = "development"
	}
	viper.SetConfigName(GoEnv)
	viper.AddConfigPath(getConfigPath())

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s \n", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("failed to unmarshal err: %s \n", err))
	}
}

const repositoryName = "gqlgen-todos"

func getConfigPath() string {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, repositoryName) && !strings.HasSuffix(wd, "app") {
		wd = filepath.Dir(wd)
	}
	if wd == "/app" {
		return wd + "/src/config"
	}

	return wd + "/api/src/config"
}
