package db

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"io/fs"
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

// extractEmbeddedDb copies the embedded seed database to the target path.
// It only writes the file if it does not already exist.
func extractEmbeddedDb(embeddedFS embed.FS, embeddedName string, destPath string) error {
	if doesItExist(destPath) {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
		return fmt.Errorf("create db directory: %w", err)
	}

	data, err := fs.ReadFile(embeddedFS, embeddedName)
	if err != nil {
		return fmt.Errorf("read embedded db: %w", err)
	}

	if err := os.WriteFile(destPath, data, 0o644); err != nil {
		return fmt.Errorf("write db to disk: %w", err)
	}

	return nil
}

// AppDbPath returns the path where the app database should live, inside the
// user's OS config directory (e.g. ~/.config/jonah/DRC.db on Linux/macOS,
// %AppData%/jonah/DRC.db on Windows).
func AppDbPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("get user config dir: %w", err)
	}

	return filepath.Join(configDir, "jonah", "DRC.db"), nil
}

// SetupEmbeddedDb extracts the embedded seed database to the user's config
// directory (if it doesn't already exist), opens it, configures pragmas, and
// optionally runs migrations.
func SetupEmbeddedDb(embeddedFS embed.FS, embeddedName string, doMigrations bool) *sql.DB {
	dbPath, err := AppDbPath()
	if err != nil {
		log.Fatal("unable to determine db path: ", err)
	}

	if err := extractEmbeddedDb(embeddedFS, embeddedName, dbPath); err != nil {
		log.Fatal("unable to extract embedded db: ", err)
	}

	return SetupDb(dbPath, doMigrations)
}

// SetupDb opens a SQLite database at the given path, creates it if missing,
// configures pragmas, and optionally runs migrations.
func SetupDb(path string, doMigrations bool) *sql.DB {
	if exists := doesItExist(path); !exists {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			log.Fatal(err)
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
		if err != nil {
			if !errors.Is(err, os.ErrExist) {
				log.Fatal(err)
			}
		} else {
			f.Close()
		}
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
