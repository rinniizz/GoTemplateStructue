package database

import (
	"fmt"

	"go-template-structure/internal/config"
	"go-template-structure/internal/domain"
	"go-template-structure/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// NewPostgresConnection creates a new PostgreSQL database connection
func NewPostgresConnection(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)

	// Configure GORM logger
	gormLog := gormLogger.Default.LogMode(gormLogger.Info)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLog,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("Successfully connected to PostgreSQL database")

	// Auto migrate tables
	if err := autoMigrate(db); err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %w", err)
	}

	return db, nil
}

// autoMigrate performs database migration
func autoMigrate(db *gorm.DB) error {
	logger.Info("Starting database migration...")

	err := db.AutoMigrate(
		&domain.User{},
		// Add more models here
	)

	if err != nil {
		return fmt.Errorf("auto migration failed: %w", err)
	}

	logger.Info("Database migration completed successfully")
	return nil
}
