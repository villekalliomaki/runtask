package config

import (
	"github.com/caarlos0/env/v7"
	"go.uber.org/zap"
)

// All app configuration is here
type Config struct {
	// Listen port for the web server
	Port int `env:"PORT" envDefault:"8080"`
	// Host for the web server
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	// PostgreSQL connection string
	PostgresConnection string `env:"PG_CONNECTION" envDefault:"host=localhost user=runtask password=runtask dbname=runtask port=5432 sslmode=prefer"`
}

// Get the app configuration from the environment
func New(log *zap.Logger) Config {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		log.Fatal("Failed to generate config", zap.Error(err))
	}

	return cfg
}
