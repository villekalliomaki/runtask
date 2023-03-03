package db

import (
	"runtask/internal/config"
	"runtask/internal/task"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Create a new GORM database connection
func New(log *zap.Logger, conf config.Config) *gorm.DB {
	log.Info("Starting database connection")

	db, err := gorm.Open(postgres.Open(conf.PostgresConnection))

	if err != nil {
		log.Panic("Failed to start PostgreSQL connection", zap.Error(err))
	}

	migrate(log, db)

	return db
}

// Run GORM migrations
func migrate(log *zap.Logger, db *gorm.DB) {
	log.Info("Running database migrations")

	if err := db.AutoMigrate(&task.DefinitionProperties{}); err != nil {
		log.Panic("Failed to run migrations", zap.Error(err))
	}
}
