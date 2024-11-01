// config/config.go
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string

	LostarkAPIKey string
}

func GetConfig() *Config {
	once.Do(func() {
		envPath := filepath.Join("env", ".env")
		instance = loadConfig(envPath)
	})
	return instance
}

func loadConfig(envPath string) *Config {
	if err := godotenv.Load(envPath); err != nil {
		fmt.Printf("Warning: .env file not found. Using environment variables: %v\n", err)
	}

	config := &Config{
		// Database
		DBHost:     getEnvWithDefault("DB_HOST", "localhost"),
		DBPort:     getEnvWithDefault("DB_PORT", "5432"),
		DBUser:     getEnvWithDefault("DB_USER", "user"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME"),

		LostarkAPIKey: getEnv("LOSTARK_API_KEY"),
	}

	return config
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("필수 환경변수가 설정되지 않았습니다: %s", key))
	}
	return value
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
