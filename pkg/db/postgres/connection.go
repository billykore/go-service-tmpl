package postgres

import (
	"github.com/billykore/go-service-tmpl/pkg/config"
	"github.com/billykore/go-service-tmpl/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New returns new postgres db connection.
func New(cfg *config.Config) *gorm.DB {
	dsn := cfg.Postgres.DSN
	log := logger.New()
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Usecase("New").Fatalf("failed to connect database: %v", err)
		return nil
	}
	return db
}
