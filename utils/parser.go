package utils

import "github.com/akshayr96/bounceSearch/types"

// Checks if the given word is indexed in the given tree
// Takes the pointer to a ternaty tree and the word to be searched in the tree
// Returns a StringMatch Enum attribute if its a partial / complete / no match
func Parser(tree *types.Node, word string) (exists types.StringMatch) {
	for len(word) > 0 {
		if tree == nil {
			return types.NoMatch
		}
		if string(word[0]) < tree.Data {
			tree = tree.Left
			continue
		}
		if string(word[0]) > tree.Data {
			tree = tree.Right
			continue
		}
		if string(word[0]) == tree.Data {
			word = word[1:]
			if len(word) == 0 {
				if tree.End == true {
					return types.CompleteMatch
				} else {
					return types.PartialMatch
				}
			}
			tree = tree.Mid
			continue
		}
	}
	return types.NoMatch
}
