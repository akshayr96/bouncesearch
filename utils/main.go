package utils

import (
	"github.com/akshayr96/bounceSearch/types"
)

func InsertWord(tree *types.Node, word string) {
	if len(word) > 0 {
		if tree.Data == "" {
			insertIntoEmptySubtree(tree, word)
		} else {
			_insertWord(tree, word)
		}
	}
}

func insertIntoEmptySubtree(subtree *types.Node, word string) {
	if len(word) > 0 {
		for len(word) > 0 {
			if subtree.Data == "" {
				subtree.Data = string(word[0])
			} else {
				subtree.Mid = &types.Node{Data: string(word[0])}
				subtree = subtree.Mid
			}
			word = word[1:]
		}
		subtree.End = true
	}
}

func _insertWord(tree *types.Node, word string) {
	for len(word) > 0 {
		// fmt.Println(tree)
		// fmt.Println(word)
		if string(word[0]) > tree.Data {
			if tree.Right == nil {
				tree.Right = &types.Node{}
			}
			if tree.Right.Data == "" {
				insertIntoEmptySubtree(tree.Right, word)
				return
			} else {
				tree = tree.Right
			}
		} else if string(word[0]) < tree.Data {
			if tree.Left == nil {
				tree.Left = &types.Node{}
			}
			if tree.Left.Data == "" {
				insertIntoEmptySubtree(tree.Left, word)
				return
			} else {
				tree = tree.Left
			}
		} else if string(word[0]) == tree.Data {
			word = word[1:]
			if tree.Mid == nil {
				tree.Mid = &types.Node{}
			}
			if tree.Mid.Data == "" {
				insertIntoEmptySubtree(tree.Mid, word)
				return
			} else {
				if len(word) > 0 {
					tree = tree.Mid
				} else {
					tree.End = true
				}
			}
		}
	}
	return
}

func Parser(tree *types.Node, word string) (exists bool) {
	for len(word) > 0 {
		if tree == nil {
			return false
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
				// return partial or complete match based on end == true
				return true
			}
			tree = tree.Mid
			continue
		}
	}
	return false
}
