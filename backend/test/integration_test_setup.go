package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/SturlaSolheim/mediaCircleBackend/config"
	"github.com/SturlaSolheim/mediaCircleBackend/database"
	"gorm.io/gorm"
)

func SetupIntegrationTest(t *testing.T) *gorm.DB {
	os.Setenv("GO_ENV", "testing")
	
	wd, _ := os.Getwd()
	projectRoot := findProjectRoot(wd)
	if projectRoot != "" {
		os.Chdir(projectRoot)
		t.Cleanup(func() {
			os.Chdir(wd)
		})
	}
	
	err := config.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load test config: %v", err)
	}

	err = database.InitDB()
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}

	err = database.AutoMigrate()
	if err != nil {
		t.Fatalf("Failed to run test migrations: %v", err)
	}

	t.Cleanup(func() {
		database.CloseDB()
	})

	return database.DB
}

func findProjectRoot(startDir string) string {
	dir := startDir
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return ""
}

func ClearDatabase(db *gorm.DB) {
	var tables []string
	db.Raw("SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'").Scan(&tables)
	
	for _, table := range tables {
		db.Exec("DELETE FROM " + table)
	}
}
