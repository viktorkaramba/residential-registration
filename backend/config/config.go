package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App
		PostgreSQL
		Cache
	}
	App struct {
		LogLevel       string `env:"LOG_LEVEL" env-default:"info"`
		SaveLogsToFile bool   `env:"SAVE_LOGS_TO_FILE" env-default:"false"`
		TimeLocation   string `env:"TIME_LOCATION" env-default:"Europe/Kyiv"`
	}
	PostgreSQL struct {
		Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
		Port     string `env:"POSTGRES_PORT" env-default:"5435"`
		User     string `env:"POSTGRES_USER" env-default:"postgres"`
		Password string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
		Database string `env:"POSTGRES_DATABASE" env-default:"api"`
	}
	Cache struct {
		DefaultExpiration int64 `env:"CACHE_DEFAULT_EXPIRATION" env-default:"-1"`
		CleanupInterval   int64 `env:"CACHE_CLEANUP_INTERVAL" env-default:"86_400_000"` // 24 * 60 * 60 * 1000 seconds
	}
)

func Init() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal("failed to read env", err)
	}

	return &cfg, nil
}
