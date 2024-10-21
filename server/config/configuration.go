package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	Server   *ServerConfig
	Database *DatabaseConfig
}

type ServerConfig struct {
	AllowedOrigins string `env:"ALLOWED_ORIGINS,required"`
	Port           string `env:"SERVER_PORT,default=5000"`
	Mode           string `env:"SERVER_MODE,default=release"`
}

type DatabaseConfig struct {
	URL string `env:"DATABASE_URL_LOCAL,required"`
}

func GetAppConfig(ctx context.Context) *AppConfig {
	var cfg AppConfig
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatal("error loading env", err)
	}

	return &cfg
}
