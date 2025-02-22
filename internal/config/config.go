package config

import (
	"fmt"
	"iskra/centralized/internal/helpers"
	"log"
	"path/filepath"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	APPLICATION_STATUS string `env:"APPLICATION_STATUS,required"`
	FRONTEND_URL       string `env:"FRONTEND_URL,required"`
	BACKEND_URL        string `env:"BACKEND_URL,required"`
	JWT_SECRET         string `env:"JWT_SECRET,required"`
}

var (
	cfg  *Config
	once sync.Once
)

func Initialize() *Config {
	once.Do(func() {
		envPath := filepath.Join(helpers.GetProjectRoot(), ".env")

		fmt.Println(envPath)

		if err := godotenv.Load(envPath); err != nil {
			log.Fatal("No .env file found!")
		}

		cfg = &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	})

	return cfg
}
