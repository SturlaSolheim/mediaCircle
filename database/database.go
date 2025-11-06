package database

import (
	"log"

	"github.com/SturlaSolheim/mediaCircleBackend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitInMemoryDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("sqlite in memory database kjører")
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
