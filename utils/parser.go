package utils

import "github.com/akshayr96/bounceSearch/types"

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
