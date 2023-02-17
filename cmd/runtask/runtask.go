package main

import (
	"runtask/internal/config"
	"runtask/internal/logs"
)

func main() {
	config.GenerateConfig()

	logs.Logger.Info("Starting application")
}
