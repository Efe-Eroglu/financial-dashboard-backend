package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	JWTSecret  string
}

var AppConfig Config

func LoadConfig() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	err := godotenv.Load("config/environment/" + env + ".env")
	if err != nil {
		log.Printf("No specific env file found for %s, using default .env", env)
		godotenv.Load()
	}

	AppConfig = Config{
		ServerPort: getEnvOrDefault("SERVER_PORT", "8080"),
		DBHost:     getEnvOrPanic("DB_HOST"),
		DBPort:     getEnvOrDefault("DB_PORT", "5432"),
		DBUser:     getEnvOrPanic("DB_USER"),
		DBPassword: getEnvOrPanic("DB_PASSWORD"),
		DBName:     getEnvOrPanic("DB_NAME"),
		DBSSLMode:  getEnvOrDefault("DB_SSLMODE", "disable"),
		// JWTSecret:  getEnvOrPanic("JWT_SECRET"),
	}
}

func getEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Required environment variable %s is not set", key)
	}
	return value
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
