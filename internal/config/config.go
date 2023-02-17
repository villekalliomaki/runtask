package config

import (
	"os"
	"runtask/internal/logs"

	"github.com/cristalhq/aconfig"
	"go.uber.org/zap"
)

// Cental configuration for the application
type Config struct {
	ListenPort int `default:"8080" usage:"Listen port for the API"`
	Auth       struct {
		Token string `usage:"Strong password for access control"`
	}
}

// Generate a configuration from envirtonment and flags
func GenerateConfig() Config {
	var config Config

	loader := aconfig.LoaderFor(&config, aconfig.Config{
		EnvPrefix: "RUNTASK",
		Files:     []string{"./config.yml", "/etc/runtask/config.yml"},
	})

	if err := loader.Load(); err != nil {
		logs.Logger.Error("Failed to generate config", zap.Error(err))
		os.Exit(1)
	}

	return config
}
