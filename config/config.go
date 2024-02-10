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
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error!!!", err)
	}

	cfg := Config{}

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "5432"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "sevinch"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1218"))
	cfg.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DB", "storesell"))

	//cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost1"))
	//cfg.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "5432"))
	//cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	//cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "your password"))
	//cfg.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DB", "your db"))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}
