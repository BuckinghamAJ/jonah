package reference

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	drcBible "github.com/BuckinghamAJ/jonah/drcBible/dto"
)

type BibleReference struct {
	Book       string
	BookId     *int64
	Chapter    uint8
	StartVerse uint8
	EndVerse   uint8
	Text       string
}

func (br *BibleReference) getBookId(ctx context.Context, queries *drcBible.Queries) int {

	book, err := queries.GetBookFromTitle(ctx, sql.NullString{String: br.Book, Valid: true})

	if err != nil {
		log.Fatal("Could not find Bible Book: " + br.Book)
	}

	br.BookId = &book.ID

	return int(book.ID)
}

func (br *BibleReference) GetFullText(ctx context.Context) string {
	var app = *internal.App
	queries := app.Queries

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var bookID int

	if br.BookId == nil {
		bookID = br.getBookId(ctxWithTimeout, queries)
	} else {
		bookID = int(*br.BookId)
	}

	queryParams := drcBible.GetVersesParams{
		BookID:     sql.NullInt64{Int64: int64(bookID), Valid: true},
		Chapter:    sql.NullInt64{Int64: int64(br.Chapter), Valid: true},
		StartVerse: sql.NullInt64{Int64: int64(br.StartVerse), Valid: true},
		EndVerse:   sql.NullInt64{Int64: int64(br.EndVerse), Valid: true},
	}

	// fmt.Println(queryParams)

	verses, err := queries.GetVerses(ctxWithTimeout, queryParams)

	if err != nil {
		log.Fatal("Error Grabbing Verses: " + err.Error())
	}

	br.Text = formattingVerses(verses, br)

	return br.Text
}

func formattingVerses(Rows []drcBible.GetVersesRow, br *BibleReference) string {
	tmpVerses := []string{}

	tmpVerses = append(tmpVerses,
		fmt.Sprintf("# %s %d:%d-%d", br.Book, br.Chapter, br.StartVerse, br.EndVerse))

	for _, verseRow := range Rows {
		tmpVerses = append(tmpVerses,
			fmt.Sprintf("%d. %s", verseRow.Verse.Int64, verseRow.Text.String))
	}

	return strings.Join(tmpVerses, "\n")
}

func NewBibleReference(book string, chapter uint8, startVerse uint8, endVerse uint8) *BibleReference {
	return &BibleReference{
		Book:       book,
		Chapter:    chapter,
		StartVerse: startVerse,
		EndVerse:   endVerse,
	}
}

func ParseInputVerses(args []string) (*BibleReference, error) {
	fmt.Printf("Args: %v\n", args)
	if len(args) == 0 {
		return johnBibleReferece()
	}

	firstArg := args[0]

	pattern := `^\s*(?P<book>(?:\d+\s*)?[A-Za-z]+(?:\s+[A-Za-z]+)*)\s*(?P<chapter>\d+)\s*:\s*(?P<startVerse>\d+)(?:\s*-\s*(?P<endVerse>\d+))?\s*$`

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(firstArg)

	if len(matches) <= 3 {
		return nil, fmt.Errorf("invalid bible reference format: %s", firstArg)
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

	return NewBibleReference(book, uint8(chapter), uint8(startVerse), uint8(endVerse)), nil
}

func ParseSingleInputVerses(verse string) (*BibleReference, error) {
	passThroughVerse := []string{verse}
	return ParseInputVerses(passThroughVerse)
}

func johnBibleReferece() (*BibleReference, error) {
	testBook := "John"
	testChapter := uint8(1)
	testVerseStart := uint8(1)
	testVerseEnd := uint8(18)
	return NewBibleReference(testBook, testChapter, testVerseStart, testVerseEnd), nil
}
