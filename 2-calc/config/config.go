package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("предупреждение: .env файл не найден")
	}

	key := os.Getenv("KEY")
	if key == "" {
		return nil, fmt.Errorf("переменная KEY не установлена")
	}

	return &Config{Key: key}, nil
}
