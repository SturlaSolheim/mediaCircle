package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitInMemoryDB() error {
	var err error
	DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("sqlite in memory database kjører")
	return nil
}

func CreateTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS album (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	
	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Tabellene er laget")
	return nil
}

func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
