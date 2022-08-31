package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DatabaseURL          string
	DatabaseRetryDelay   time.Duration
	DatabaseMaxRetries   int32
	DatabaseMigrationDir string
	LogLevel             string
	HTTPServerAddress    string
}

func NewConfig() Config {
	return Config{
		DatabaseURL:          getString("DATABASE_URL", ""),
		DatabaseRetryDelay:   time.Duration(getInt32("DATABASE_RETRY_DELAY", 5)) * time.Second,
		DatabaseMaxRetries:   getInt32("DATABASE_MAX_RETRIES", 10),
		DatabaseMigrationDir: getString("DATABASE_MIGRATION_DIR", "file://db/migrations"),
		LogLevel:             getString("LOG_LEVEL", "DEBUG"),
		HTTPServerAddress:    getString("HTTP_SERVER_ADDRESS", ":8080"),
	}
}

func getString(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return val
}

func getInt32(key string, defaultValue int32) int32 {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	result, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int32(result)
}
