package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/BuckinghamAJ/jonah/internal/db"
	drcBible "github.com/BuckinghamAJ/jonah/internal/drcBible/dto"
	"github.com/BuckinghamAJ/jonah/internal/parser"
	"github.com/BuckinghamAJ/jonah/internal/reference"
)

type BookResponse struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

type BibleService struct {
	db      *sql.DB
	queries *drcBible.Queries
}

func NewBibleService(dbConn *sql.DB, queries *drcBible.Queries) *BibleService {
	return &BibleService{
		db:      dbConn,
		queries: queries,
	}
}

func (s *BibleService) ensureReady() error {
	if s.db == nil || s.queries == nil {
		return errors.New("Database is unavailable")
	}

	if db.IsBusy(s.db) {
		return errors.New("Database is busy")
	}

	return nil
}

func (s *BibleService) SearchVerse(ctx context.Context, passages string) (*reference.BibleReference, error) {
	bibleRef := parser.BiblePassageParser(passages)

	if err := s.ensureReady(); err != nil {
		if s.db == nil || s.queries == nil {
			return nil, errors.New("Database is busy with migrations")
		}
		return nil, err
	}

	bibleRef.LoadAllText(ctx, s.queries)

	return &bibleRef, nil
}

func (s *BibleService) GetAllBooks(ctx context.Context) ([]BookResponse, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}

	books, err := s.queries.GetAllBooks(ctx)
	if err != nil {
		return nil, err
	}

	bookNames := make([]BookResponse, 0, len(books))
	for _, book := range books {
		bookNames = append(bookNames, BookResponse{
			Name: book.Name.String,
			ID:   book.ID,
		})
	}

	return bookNames, nil
}

func (s *BibleService) GetAllChapters(ctx context.Context, book int64) ([]int64, error) {
	if err := s.ensureReady(); err != nil {
		return nil, err
	}

	bookIDParam := sql.NullInt64{Int64: book, Valid: true}

	chapters, err := s.queries.GetChaptersOfBook(ctx, bookIDParam)
	if err != nil {
		return nil, err
	}

	chapterNumbers := make([]int64, 0, len(chapters))
	for _, chapter := range chapters {
		if chapter.Valid {
			chapterNumbers = append(chapterNumbers, chapter.Int64)
		}
	}

	return chapterNumbers, nil
}
