package logs

import (
	"log"

	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("Failed to start zap logger: %v", err)
	}

	return logger
}
