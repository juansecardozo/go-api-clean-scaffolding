package bootstrap

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"time"
)

type Config struct {
	Host            string        `env:"APP_HOST,required"`
	Port            uint          `env:"APP_PORT,required"`
	ShutdownTimeout time.Duration `env:"APP_SHUTDOWN_TIMEOUT,required"`
}

func LoadConfig() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{}

	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
