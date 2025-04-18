package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	Error       = "error >>> "
	Info        = "info >>> "
	Log         = "log >>> "
	SecretUser  = "awe-1234"
	SecretAdmin = "sec-1234"
)

type Config struct {
	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string

	ServiceHost     string
	ServiceHTTPPort string
}

func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Println("not found env")
	}

	var cfg Config

	cfg.ServiceHost = cast.ToString(getValueOrDefault("SERVICE_HOST", "localhost"))
	cfg.ServiceHTTPPort = cast.ToString(getValueOrDefault("SERVICE_HTTP_PORT", ":8080"))

	cfg.PostgresHost = cast.ToString(getValueOrDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresUser = cast.ToString(getValueOrDefault("POSTGRES_USER", "postgres"))
	cfg.PostgresDatabase = cast.ToString(getValueOrDefault("POSTGRES_DATABASE", "travel"))
	cfg.PostgresPassword = cast.ToString(getValueOrDefault("POSTGRES_PASSWORD", "1234"))
	cfg.PostgresPort = cast.ToString(getValueOrDefault("POSTGRES_PORT", "5432"))
	// cfg.Password = cast.ToString(getValueOrDefault("PASSWORD", "1234"))

	return cfg
}

func getValueOrDefault(key string, defaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
