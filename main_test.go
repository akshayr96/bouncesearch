package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/akshayr96/bounceSearch/core"
	"github.com/akshayr96/bounceSearch/provider"
	"github.com/akshayr96/bounceSearch/types"
)

// Tests the indexer and the parser
func TestIndexerAndParser(t *testing.T) {
	tree := &types.Node{}
	testWords := []string{"plumbus", "eyeholeman", "morty", "glibglob", "squanchy"}

	//Indexing words
	for i, word := range testWords {
		core.InsertWord(tree, word, strconv.Itoa(i), 1.0)
	}

	//writes the tree file for reference
	provider.WriteTree(tree, "randomapp", "users")

	//tests for complete match
	for _, word := range testWords {
		if core.Parser(tree, word) != types.CompleteMatch {
			t.Errorf("Complete match test failed for the word  %s", word)
		}
	}

	fmt.Println("PASS: Complete Match test")

	//tests for partial match
	partialMatchWords := []string{"plumb", "eye", "mort", "glib", "squanch"}
	for _, word := range partialMatchWords {
		if core.Parser(tree, word) != types.PartialMatch {
			t.Errorf("Partial match test failed for the word %s", word)
		}
	}

	fmt.Println("PASS: Partial Match test")

	//tests for no match
	noMatchWords := []string{"froopyland", "summer", "birdperson", "fakedoors", "shmoopydoop"}
	for _, word := range noMatchWords {
		if core.Parser(tree, word) != types.NoMatch {
			t.Errorf("No match test failed for the word %s", word)
		}
	}

	fmt.Println("PASS: No Match test")

}

func TestLevenshtein(t *testing.T) {
	if core.LevenshteinDistance("app", "map") != 2 {
		t.Error("Levenshtein test failed")
	}
	fmt.Println("PASS: Levenshtein Test")
}
