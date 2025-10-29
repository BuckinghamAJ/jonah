package reference

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
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

func (br *BibleReference) ExtractBiblePassages(input []string) error {
	for _, v := range input {
		passage, err := ExtractPassageInfo(v)
		if err != nil {
			return err
		}
		br.Passages = append(br.Passages, passage)
	}

	return nil
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
	Text       string
}

func (bp *BiblePassage) getBookId(ctx context.Context, queries *drcBible.Queries) int {

	book, err := queries.GetBookFromTitle(ctx, sql.NullString{String: bp.Book, Valid: true})

	if err != nil {
		log.Fatal("Could not find Bible Book: " + bp.Book)
	}

	bp.BookId = &book.ID

	return int(book.ID)
}

func (bp *BiblePassage) GetFullText(ctx context.Context, queries *drcBible.Queries) string {

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var bookID int

	if bp.BookId == nil {
		bookID = bp.getBookId(ctxWithTimeout, queries)
	} else {
		bookID = int(*bp.BookId)
	}

	queryParams := drcBible.GetVersesParams{
		BookID:     sql.NullInt64{Int64: int64(bookID), Valid: true},
		Chapter:    sql.NullInt64{Int64: int64(bp.Chapter), Valid: true},
		StartVerse: sql.NullInt64{Int64: int64(bp.StartVerse), Valid: true},
		EndVerse:   sql.NullInt64{Int64: int64(bp.EndVerse), Valid: true},
	}

	// fmt.Println(queryParams)

	verses, err := queries.GetVerses(ctxWithTimeout, queryParams)

	if err != nil {
		log.Fatal("Error Grabbing Verses: " + err.Error())
	}

	bp.Text = formattingVerses(verses, bp)

	return bp.Text
}

func formattingVerses(Rows []drcBible.GetVersesRow, bp *BiblePassage) string {
	tmpVerses := []string{}

	tmpVerses = append(tmpVerses,
		fmt.Sprintf("# %s %d:%d-%d", bp.Book, bp.Chapter, bp.StartVerse, bp.EndVerse))

	for _, verseRow := range Rows {
		tmpVerses = append(tmpVerses,
			fmt.Sprintf("%d. %s", verseRow.Verse.Int64, verseRow.Text.String))
	}

	return strings.Join(tmpVerses, "\n")
}

func NewBiblePassage(book string, chapter uint8, startVerse uint8, endVerse uint8) *BiblePassage {
	return &BiblePassage{
		Book:       book,
		Chapter:    chapter,
		StartVerse: startVerse,
		EndVerse:   endVerse,
	}
}

func ExtractPassageInfo(passage string) (*BiblePassage, error) {
	fmt.Printf("passage: %v\n", passage)

	pattern := `^\s*(?P<book>(?:\d+\s*)?[A-Za-z]+(?:\s+[A-Za-z]+)*)\s*(?P<chapter>\d+)\s*:\s*(?P<startVerse>\d+)(?:\s*-\s*(?P<endVerse>\d+))?\s*$`

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(passage)

	if len(matches) <= 3 {
		return nil, fmt.Errorf("invalid bible reference format: %s", passage)
	}

	book := matches[1]

	chapter, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, fmt.Errorf("invalid chapter number: %s", matches[2])
	}

	startVerse, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, fmt.Errorf("invalid start verse: %s", matches[3])
	}

	endVerse, err := strconv.Atoi(matches[4])
	if err != nil && startVerse == 0 {
		return nil, fmt.Errorf("invalid end verse: %s", matches[4])
	}
	if endVerse == 0 {
		endVerse = startVerse
	}

	return NewBiblePassage(book, uint8(chapter), uint8(startVerse), uint8(endVerse)), nil
}
