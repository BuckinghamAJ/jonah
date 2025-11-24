package parser

import (
	"reflect"
	"testing"

	"github.com/BuckinghamAJ/jonah/reference"
)

func TestBibleVerseParser(t *testing.T) {

	expectedBibleRef := reference.BibleReference{
		Passages: []*reference.BiblePassage{
			{
				Book:       "John",
				Chapter:    uint8(3),
				StartVerse: uint8(1),
				EndVerse:   uint8(16),
			},
			{
				Book:       "Psalms",
				Chapter:    uint8(8),
				StartVerse: uint8(1),
				EndVerse:   uint8(15),
			},
		},
	}

	got := BiblePassageParser("John 3:1-16; Psalms 8:1-15")

	if !reflect.DeepEqual(got, expectedBibleRef) {

		t.Fatalf("mismatch: %#v vs %#v", got, expectedBibleRef)
	}

}

func TestBibleVerseParserWithVerseandVerseRange(t *testing.T) {

	expectedBibleRef := reference.BibleReference{
		Passages: []*reference.BiblePassage{
			{
				Book:       "John",
				Chapter:    uint8(3),
				StartVerse: uint8(1),
			},
			{
				Book:       "Psalms",
				Chapter:    uint8(8),
				StartVerse: uint8(1),
				EndVerse:   uint8(15),
			},
		},
	}

	got := BiblePassageParser("John 3:1; Psalms 8:1-15")

	if !reflect.DeepEqual(got, expectedBibleRef) {

		t.Fatalf("mismatch: %#v vs %#v", got, expectedBibleRef)
	}

}

func TestBibleVerseParserWithWholeChapterandVerseRange(t *testing.T) {

	expectedBibleRef := reference.BibleReference{
		Passages: []*reference.BiblePassage{
			{
				Book:    "John",
				Chapter: uint8(3),
			},
			{
				Book:       "Psalms",
				Chapter:    uint8(8),
				StartVerse: uint8(1),
				EndVerse:   uint8(15),
			},
		},
	}

	got := BiblePassageParser("John 3; Psalms 8:1-15")

	if !reflect.DeepEqual(got, expectedBibleRef) {

		t.Fatalf("mismatch: %#v vs %#v", got, expectedBibleRef)
	}

}

func TestBibleVerseParserWithSingleVerseRange(t *testing.T) {

	expectedBibleRef := reference.BibleReference{
		Passages: []*reference.BiblePassage{
			{
				Book:       "Psalms",
				Chapter:    uint8(8),
				StartVerse: uint8(1),
				EndVerse:   uint8(15),
			},
		},
	}

	got := BiblePassageParser("Psalms 8:1-15")

	if !reflect.DeepEqual(got, expectedBibleRef) {

		t.Fatalf("mismatch: %#v vs %#v", got, expectedBibleRef)
	}

}

func TestBibleVerseParserWithSingleVerse(t *testing.T) {

	expectedBibleRef := reference.BibleReference{
		Passages: []*reference.BiblePassage{
			{
				Book:       "Psalms",
				Chapter:    uint8(8),
				StartVerse: uint8(1),
			},
		},
	}

	got := BiblePassageParser("Psalms 8:1")

	if !reflect.DeepEqual(got, expectedBibleRef) {

		t.Fatalf("mismatch: %#v vs %#v", got, expectedBibleRef)
	}

}

func TestBibleVerseParserWithWholeChapter(t *testing.T) {

	expectedBibleRef := reference.BibleReference{
		Passages: []*reference.BiblePassage{
			{
				Book:    "Psalms",
				Chapter: uint8(8),
			},
		},
	}

	got := BiblePassageParser("Psalms 8")

	if !reflect.DeepEqual(got, expectedBibleRef) {

		t.Fatalf("mismatch: %#v vs %#v", got, expectedBibleRef)
	}

}
