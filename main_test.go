package main

import (
	"testing"

	"github.com/akshayr96/bounceSearch/provider"
	"github.com/akshayr96/bounceSearch/types"
	"github.com/akshayr96/bounceSearch/utils"
)

// Tests the indexer and the parser
func TestIndexerAndParser(t *testing.T) {
	tree := &types.Node{}
	testWords := []string{"plumbus", "eyeholeman", "morty", "glibglob", "squanchy"}

	//Indexing words
	for _, word := range testWords {
		utils.InsertWord(tree, word, "1", 1.0)
	}

	//writes the tree file for reference
	provider.WriteTree(tree, "randomapp", "users")

	//tests for complete match
	for _, word := range testWords {
		if utils.Parser(tree, word) != types.CompleteMatch {
			t.Errorf("Complete match test failed for the word  %s", word)
		}
	}

	//tests for partial match
	partialMatchWords := []string{"plumb", "eye", "mort", "glib", "squanch"}
	for _, word := range partialMatchWords {
		if utils.Parser(tree, word) != types.PartialMatch {
			t.Errorf("Partial match test failed for the word %s", word)
		}
	}

	//tests for no match
	noMatchWords := []string{"froopyland", "summer", "birdperson", "fakedoors", "shmoopydoop"}
	for _, word := range noMatchWords {
		if utils.Parser(tree, word) != types.NoMatch {
			t.Errorf("No match test failed for the word %s", word)
		}
	}

}

func TestLevenshtein(t *testing.T) {
	if utils.LevenshteinDistance("app", "map") != 2 {
		t.Error("Levenshtein test failed")
	}
}
