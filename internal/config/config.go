package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	FRONTEND_URL string `env:"FRONTEND_URL,required"`
	BACKEND_URL  string `env:"BACKEND_URL,required"`
	JWT_SECRET   string `env:"JWT_SECRET,required"`
}

var (
	cfg  *Config
	once sync.Once
)

func Initialize() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("No .env file found!")
		}

		cfg = &Config{}
		if err := env.Parse(cfg); err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	})

	return cfg
}
