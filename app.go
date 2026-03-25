package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/BuckinghamAJ/jonah/internal/db"
	drcBible "github.com/BuckinghamAJ/jonah/internal/drcBible/dto"
	"github.com/BuckinghamAJ/jonah/internal/reference"
	"github.com/BuckinghamAJ/jonah/internal/services"
	_ "github.com/mattn/go-sqlite3"
)

// App struct
type App struct {
	ctx          context.Context
	database     *sql.DB
	bibleService *services.BibleService
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

	a.database = db.SetupDb(dbPath, true)

	a.bibleService = services.NewBibleService(
		a.database,
		drcBible.New(a.database),
	)

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
	if err := a.database.Close(); err != nil {
		log.Println("unable to close database:", err)
	}
}

func (a *App) SearchVerse(passages string) (*reference.BibleReference, error) {
	return a.bibleService.SearchVerse(a.ctx, passages)
}

func (a *App) GetAllBooks() ([]services.BookResponse, error) {
	return a.bibleService.GetAllBooks(a.ctx)
}

func (a *App) GetAllChapters(book int64) ([]int64, error) {
	return a.bibleService.GetAllChapters(a.ctx, book)
}
