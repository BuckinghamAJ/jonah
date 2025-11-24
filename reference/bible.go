package reference

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	drcBible "github.com/BuckinghamAJ/jonah/drcBible/dto"
)

type BibleReference struct {
	Passages []*BiblePassage
}

func NewBibleReference() *BibleReference {
	return &BibleReference{
		Passages: []*BiblePassage{},
	}
}

func (br *BibleReference) LoadAllText(ctx context.Context, queries *drcBible.Queries) {
	var wg sync.WaitGroup
	for i := range br.Passages {
		wg.Add(1)
		go func(p *BiblePassage) {
			defer wg.Done()
			p.GetFullText(ctx, queries)
		}(br.Passages[i])
	}
	wg.Wait()
}

type BiblePassage struct {
	Book       string
	BookId     *int64
	Chapter    uint8
	StartVerse uint8
	EndVerse   uint8
	FullText   []Verse
}

type Verse struct {
	Number int64
	Text   string
}

func (bp *BiblePassage) getBookId(ctx context.Context, queries *drcBible.Queries) (int, error) {

	book, err := queries.GetBookFromTitle(ctx, sql.NullString{String: bp.Book, Valid: true})

	if err != nil {
		return 0, fmt.Errorf("Could not find Bible Book: %s", bp.Book)
	}

	bp.BookId = &book.ID

	return int(book.ID), nil
}

func (bp *BiblePassage) GetFullText(ctx context.Context, queries *drcBible.Queries) []Verse {

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var bookID int
	var verses any
	var err error

	if bp.BookId == nil {
		bookID, err = bp.getBookId(ctxWithTimeout, queries)
	} else {
		bookID = int(*bp.BookId)
	}

	if err != nil {
		fmt.Println("Error in GetFullText:", err)
		return make([]Verse, 0)
	}

	if bp.StartVerse == 0 {
		queryParams := drcBible.MakeChapterParams(bookID, bp.Chapter)
		verses, err = queries.GetChapter(ctxWithTimeout, queryParams)
	} else {
		queryParams := drcBible.MakeVerseParams(bookID, bp.Chapter, bp.StartVerse, bp.EndVerse)
		verses, err = queries.GetVerses(ctxWithTimeout, queryParams)
	}

	if err != nil {
		fmt.Println("Error in GetFullText:", err)
		return make([]Verse, 0)
	}

	bp.FullText = formattingVerses(verses)
	return bp.FullText
}

func formattingVerses(verses any) []Verse {
	switch v := verses.(type) {
	case []drcBible.GetChapterRow:
		return formatChapterRows(v)
	case []drcBible.GetVersesRow:
		return formatVerseRows(v)
	default:
		return nil
	}
}

func formatChapterRows(Rows []drcBible.GetChapterRow) []Verse {
	tmpVerses := []Verse{}

	for _, verseRow := range Rows {
		tmpVerses = append(tmpVerses,
			Verse{verseRow.Verse.Int64, verseRow.Text.String})
	}

	return tmpVerses
}

func formatVerseRows(Rows []drcBible.GetVersesRow) []Verse {
	tmpVerses := []Verse{}

	for _, verseRow := range Rows {
		tmpVerses = append(tmpVerses,
			Verse{verseRow.Verse.Int64, verseRow.Text.String})
	}

	return tmpVerses
}

func NewBiblePassage(book string, chapter uint8, startVerse uint8, endVerse uint8) *BiblePassage {
	return &BiblePassage{
		Book:       book,
		Chapter:    chapter,
		StartVerse: startVerse,
		EndVerse:   endVerse,
	}
}
