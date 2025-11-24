package parser

import (
	"fmt"
	"strconv"

	"github.com/BuckinghamAJ/jonah/reference"
	parsec "github.com/prataprc/goparsec"
)

func makeBibleParser(ast *parsec.AST) parsec.Parser {
	var verse, verseRange, singleVerse, passage parsec.Parser

	// Example input John 3:1-16; Psalms 9:1-70

	toVerse := parsec.Atom(":", "TOV")
	splitVerse := parsec.Atom("-", "SPV")
	book := parsec.Token(`[a-zA-Z0-9]+`, "BOOK")
	chapter := parsec.Token(`[0-9]+`, "CHAPTER")
	verseNum := parsec.Token(`[0-9]+`, "VERSE")
	separator := parsec.Atom(";", "SEPARATOR")

	//Whole Chapter John 3
	wholeChapter := ast.And("whole_chapter", nil, book, chapter)

	// Single verse: John 3:16
	singleVerse = ast.And("single_verse", nil, book, chapter, toVerse, verseNum)

	// Verse range: John 3:1-16
	verseRange = ast.And("verse_range", nil, book, chapter, toVerse, verseNum, splitVerse, verseNum)

	// Either a single verse or a range
	verse = ast.OrdChoice("verse", nil, verseRange, singleVerse, wholeChapter)

	// Passage: multiple verses/ranges separated by ';'
	passage = ast.Kleene("passage", nil,
		ast.And("passage_item", nil, verse, ast.Maybe("sep", nil, separator)),
		nil,
	)

	return passage
}

func BiblePassageParser(input string) reference.BibleReference {
	ast := parsec.NewAST("bible", 100)
	bParser := makeBibleParser(ast)

	s := parsec.NewScanner([]byte(input))

	node, _ := ast.Parsewith(bParser, s)

	return walkAST(node)
}

func walkAST(node parsec.Queryable) reference.BibleReference {
	ref := reference.BibleReference{
		Passages: make([]*reference.BiblePassage, 0),
	}

	children := node.GetChildren()

	for _, child := range children {
		if child.GetName() == "passage_item" {
			passage := extractPassage(child)
			if passage != nil {
				ref.Passages = append(ref.Passages, passage)
			}
		}

	}

	return ref
}

func extractPassage(node parsec.Queryable) *reference.BiblePassage {
	children := node.GetChildren()
	if len(children) == 0 {
		fmt.Println("extractPassage: node has no children")
		return nil
	}

	verseNode := children[0]
	verseChildren := verseNode.GetChildren()

	if len(verseChildren) < 2 {
		return nil
	}

	book := verseChildren[0].GetValue()
	chapterStr := verseChildren[1].GetValue()
	chapter, err := strconv.ParseUint(chapterStr, 10, 8)
	if err != nil {
		return nil
	}
	fmt.Printf("extractPassage: book = %s, chapter = %d\n", book, chapter)

	switch verseNode.GetName() {
	case "whole_chapter":
		return &reference.BiblePassage{
			Book:    book,
			Chapter: uint8(chapter),
		}
	case "single_verse":
		if len(verseChildren) < 4 {
			return nil
		}
		verseStr := verseChildren[3].GetValue()
		verse, err := strconv.ParseUint(verseChildren[3].GetValue(), 10, 8)
		if err != nil {
			fmt.Printf("extractPassage: error parsing verse '%s': %v\n", verseStr, err)
			return nil
		}
		return &reference.BiblePassage{
			Book:       book,
			Chapter:    uint8(chapter),
			StartVerse: uint8(verse),
		}
	case "verse_range":
		if len(verseChildren) < 6 {
			return nil
		}
		startVerseStr := verseChildren[3].GetValue()
		endVerseStr := verseChildren[5].GetValue()
		startVerse, err1 := strconv.ParseUint(startVerseStr, 10, 8)
		endVerse, err2 := strconv.ParseUint(endVerseStr, 10, 8)
		if err1 != nil || err2 != nil {
			fmt.Printf("extractPassage: error parsing startVerse '%s' or endVerse '%s': %v, %v\n", startVerseStr, endVerseStr, err1, err2)
			return nil
		}
		return &reference.BiblePassage{
			Book:       book,
			Chapter:    uint8(chapter),
			StartVerse: uint8(startVerse),
			EndVerse:   uint8(endVerse),
		}
	default:
		fmt.Printf("extractPassage: unknown verseNode name '%s'\n", verseNode.GetName())
		return nil
	}

}
