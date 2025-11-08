package database

import (
	"log"

	"github.com/SturlaSolheim/mediaCircleBackend/config"
	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	connection := config.AppConfig.Database.Connection
	DB, err = gorm.Open(sqlite.Open(connection), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Printf("Database initialized with connection: %s", connection)
	return nil
}

func CreateTables() error {
	err := DB.AutoMigrate(&models.Album{})
	if err != nil {
		return err
	}

	log.Println("Tabellene er laget")
	return nil
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
