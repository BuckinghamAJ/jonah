package db

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/BuckinghamAJ/jonah/migrations"
	"github.com/mattn/go-sqlite3"
)

func doesItExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || !errors.Is(err, os.ErrNotExist)
}

// Helper to create a SQLite DB if does not exist.
func createDb(path string) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)

	if err != nil {
		if errors.Is(err, os.ErrExist) {
			log.Println("file already exists")
			return
		}
		log.Fatal(err)
	}
	defer f.Close()

}

// Helper to setup SQLite connection
func SetupDb(path string, doMigrations bool) *sql.DB {
	if exists := doesItExist(path); !exists {
		createDb(path)
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal("unable to connect to database: ", err)
	}

	if _, err := db.Exec("PRAGMA journal_mode=WAL;"); err != nil {
		log.Fatal("unable to change into wal mode: ", err)
	}
	if _, err := db.Exec("PRAGMA busy_timeout = 5000;"); err != nil {
		log.Fatal("unable to set busy timeout: ", err)
	}

	if doMigrations {
		if err := migrations.Run(db); err != nil {
			log.Fatal("unable to run migrations: ", err)
		}
	}

	return db
}

func IsBusy(db *sql.DB) bool {
	_, err := db.Exec("SELECT 1;")
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.Code == sqlite3.ErrBusy {
			return true
		}
	}

	return false
}
