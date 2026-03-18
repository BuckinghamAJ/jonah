package drcBible

import "database/sql"

// BibleRows abstracts result row types that contain chapter, verse, and text fields.
type BibleRows interface {
	GetChapter() sql.NullInt64
	GetVerse() sql.NullInt64
	GetText() sql.NullString
}

// MakeChapterParams creates a GetChapterParams from simple Go types.
func MakeChapterParams(bookID int, chapter uint8) GetChapterParams {
	return GetChapterParams{
		BookID:  sql.NullInt64{Int64: int64(bookID), Valid: true},
		Chapter: sql.NullInt64{Int64: int64(chapter), Valid: true},
	}
}

// Method implementations for GetChapterRow to satisfy BibleRows.
func (p GetChapterRow) GetChapter() sql.NullInt64 { return p.Chapter }
func (p GetChapterRow) GetVerse() sql.NullInt64   { return p.Verse }
func (p GetChapterRow) GetText() sql.NullString   { return p.Text }

// MakeVerseParams creates a GetVersesParams from simple Go types.
func MakeVerseParams(bookID int, chapter uint8, startVerse uint8, endVerse uint8) GetVersesParams {
	return GetVersesParams{
		BookID:     sql.NullInt64{Int64: int64(bookID), Valid: true},
		Chapter:    sql.NullInt64{Int64: int64(chapter), Valid: true},
		StartVerse: sql.NullInt64{Int64: int64(startVerse), Valid: true},
		EndVerse:   sql.NullInt64{Int64: int64(endVerse), Valid: true},
	}
}

// Method implementations for GetVersesRow to satisfy BibleRows.
func (p GetVersesRow) GetChapter() sql.NullInt64 { return p.Chapter }
func (p GetVersesRow) GetVerse() sql.NullInt64   { return p.Verse }
func (p GetVersesRow) GetText() sql.NullString   { return p.Text }
