package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TelegramToken string
}

func LoadConfig() (*Config, error) {
	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
	}

	return cfg, nil
}
