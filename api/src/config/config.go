package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/Sei-Yukinari/gqlgen-todos/src/path"
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
	viper.AddConfigPath(path.GetProjectRootPath() + "/src/config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s \n", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("failed to unmarshal err: %s \n", err))
	}
}
