package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	DebugMode  bool   `env:"DEBUG"`
	ServerHost string `env:"SERVER_HOST" envDefault:":8080"`
}

func Read() Config {
	cfg := ServerConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return Config{
		Server: cfg,
	}
}
