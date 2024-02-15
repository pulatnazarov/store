package config

import (
	"fmt"
	"github.com/spf13/cast"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string

	ServiceName string
	LoggerLevel string

	RedisHost     string
	RedisPort     string
	RedisPassword string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error!!!", err)
	}

	cfg := Config{}

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "5432"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "user"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "password"))
	cfg.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DB", "password"))

	cfg.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", "store"))
	cfg.LoggerLevel = cast.ToString(getOrReturnDefault("LOGGER_LEVEL", "debug"))

	cfg.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	cfg.RedisPort = cast.ToString(getOrReturnDefault("REDIS_PORT", "6379"))
	cfg.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", "password"))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}
