package config

import (
	"log"
	"time"
	_ "time/tzdata"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App
		Server
		Token
		PostgreSQL
		Cache
	}
	App struct {
		LogLevel       string `env:"LOG_LEVEL" env-default:"info"`
		SaveLogsToFile bool   `env:"SAVE_LOGS_TO_FILE" env-default:"false"`
		TimeLocation   string `env:"TIME_LOCATION" env-default:"Europe/Kyiv"`
		Salt           string `env:"SALT" env-default:""`
	}
	Server struct {
		Port string `env:"PORT" env-default"8080"`
	}
	Token struct {
		TokenTLL  time.Duration `env:"TOKEN_TLL" env-default:"12h"`
		SignInKey string        `env:"SIGNINKEY" env-default:""`
	}
	PostgreSQL struct {
		Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
		Port     string `env:"POSTGRES_PORT" env-default:"5435"`
		User     string `env:"POSTGRES_USER" env-default:"postgres"`
		Password string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
		Database string `env:"POSTGRES_DATABASE" env-default:"residential_registration"`
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
