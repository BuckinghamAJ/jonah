package reference

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	drcBible "github.com/BuckinghamAJ/jonah/drcBible/dto"
)

type BibleReference struct {
	Passages []*BiblePassage
}

func (br *BibleReference) ToString() string {
	var sb strings.Builder
	for i, passage := range br.Passages {
		if i > 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(passage.ToString())
	}
	return sb.String()
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

	log.Println(br.Passages)
}

type BiblePassage struct {
	Book       string
	BookId     *int64
	Chapter    uint8
	StartVerse uint8
	EndVerse   uint8
	FullText   []Verse
}

func (bp *BiblePassage) ToString() string {
	if bp.StartVerse == 0 && bp.EndVerse == 0 {
		return fmt.Sprintf("%s %d", bp.Book, bp.Chapter)
	}
	if bp.EndVerse == 0 || bp.EndVerse == bp.StartVerse {
		return fmt.Sprintf("%s %d:%d", bp.Book, bp.Chapter, bp.StartVerse)
	}
	return fmt.Sprintf("%s %d:%d-%d", bp.Book, bp.Chapter, bp.StartVerse, bp.EndVerse)
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
	var err error

	if bp.BookId == nil {
		bookID, err = bp.getBookId(ctxWithTimeout, queries)
	} else {
		bookID = int(*bp.BookId)
	}

	if err != nil {
		log.Println("error in GetFullText:", err)
		return make([]Verse, 0)
	}

	if bp.StartVerse == 0 {
		queryParams := drcBible.MakeChapterParams(bookID, bp.Chapter)
		rows, err := queries.GetChapter(ctxWithTimeout, queryParams)
		if err != nil {
			log.Println("error in GetFullText:", err)
			return make([]Verse, 0)
		}
		bp.FullText = formatRows(rows)
	} else {
		queryParams := drcBible.MakeVerseParams(bookID, bp.Chapter, bp.StartVerse, bp.EndVerse)
		rows, err := queries.GetVerses(ctxWithTimeout, queryParams)
		if err != nil {
			log.Println("error in GetFullText:", err)
			return make([]Verse, 0)
		}
		bp.FullText = formatRows(rows)
	}

	return bp.FullText
}

// formatRows converts a slice of BibleRows-compatible types to []Verse.
func formatRows[T drcBible.BibleRows](rows []T) []Verse {
	verses := make([]Verse, 0, len(rows))
	for _, r := range rows {
		verses = append(verses, Verse{
			Number: r.GetVerse().Int64,
			Text:   r.GetText().String,
		})
	}
	return verses
}
