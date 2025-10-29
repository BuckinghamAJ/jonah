package main

import (
	"context"
	"fmt"
	"strings"

	drcBible "github.com/BuckinghamAJ/jonah/drcBible/dto"
	"github.com/BuckinghamAJ/jonah/reference"
)

// App struct
type App struct {
	ctx     context.Context
	Queries *drcBible.Queries
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
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
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SearchVerse(passages string) (*reference.BibleReference, error) {
	passageSplit := strings.Split(strings.TrimSpace(passages), ";")
	bibleRef := reference.NewBibleReference()

	err := bibleRef.ExtractBiblePassages(passageSplit)

	if err != nil {
		return nil, err
	}

	bibleRef.LoadAllText(a.ctx, a.Queries)

	return bibleRef, nil

}
