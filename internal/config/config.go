package config

import (
	"os"
)

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

type Config struct {
	DB         DBConfig
	ServerPort string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func New() (*Config, error) {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	serverPort := getEnv("SERVER_PORT", "3081")

	config := Config{
		DB:         dbConfig,
		ServerPort: serverPort,
	}
	return &config, nil
}
