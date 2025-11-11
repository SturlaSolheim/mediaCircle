package database

import (
	"fmt"
	"log"

	"github.com/SturlaSolheim/mediaCircleBackend/config"
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	dbType := config.AppConfig.Database.Type
	connection := config.AppConfig.Database.Connection

	switch dbType {
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(connection), &gorm.Config{})
	case "postgres", "postgresql":
		DB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	default:
		return fmt.Errorf("unsupported database type: %s", dbType)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to %s database: %w", dbType, err)
	}

	log.Printf("Database initialized with type: %s, connection: %s", dbType, connection)
	return nil
}

func AutoMigrate() error {
	models := []interface{}{
		&models.Album{},
	}

	err := DB.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("failed to auto migrate: %w", err)
	}

	log.Printf("Auto migration completed for %d models", len(models))
	return nil
}

func CreateTables() error {
	return AutoMigrate()
}

func CloseDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
