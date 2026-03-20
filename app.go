package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/BuckinghamAJ/jonah/internal/db"
	drcBible "github.com/BuckinghamAJ/jonah/internal/drcBible/dto"
	"github.com/BuckinghamAJ/jonah/internal/parser"
	"github.com/BuckinghamAJ/jonah/internal/reference"
	_ "github.com/mattn/go-sqlite3"
)

// App struct
type App struct {
	ctx     context.Context
	queries *drcBible.Queries
	db      *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := cwd + "/data/DRC.db" //TODO: Adjust where this gets placed for install.

	a.db = db.SetupDb(dbPath, true)

	a.queries = drcBible.New(a.db)

}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	if err := a.db.Close(); err != nil {
		log.Println("unable to close database:", err)
	}
}

func (a *App) SearchVerse(passages string) (*reference.BibleReference, error) {
	bibleRef := parser.BiblePassageParser(passages)

	if a.db == nil || a.queries == nil {
		return nil, errors.New("Database is busy with migrations")
	}

	if db.IsBusy(a.db) {
		return nil, errors.New("Database is busy")
	}

	bibleRef.LoadAllText(a.ctx, a.queries)

	return &bibleRef, nil

}
