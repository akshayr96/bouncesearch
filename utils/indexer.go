package utils

import (
	"github.com/akshayr96/bounceSearch/types"
)

// Acts as a wrapper to the insertWord method to handle initial empty tree case
// Takes the pointer to a ternaty tree and the word to be indexed in the tree
func InsertWord(tree *types.Node, word string, id string, weight float32) {
	if len(word) > 0 {
		if tree.Data == "" {
			insertIntoEmptySubtree(tree, word, id, weight)
		} else {
			insertWord(tree, word, id, weight)
		}
	}
}

// Indexed data into an empty subtree
// Takes the pointer to a ternaty tree and the word to be indexed in the tree
func insertIntoEmptySubtree(subtree *types.Node, word string, id string, weight float32) {
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
		subtree.Occurances = make(types.NodeOccurances)
		weights := []float32{weight}
		subtree.Occurances[id] = types.NodeOccurance{Index: id, Weights: weights, Frequency: 1}
	}
}

//Inserts a word into the given tree pointer
// Takes the pointer to a ternaty tree and the word to be indexed in the tree
func insertWord(tree *types.Node, word string, id string, weight float32) {
	for len(word) > 0 {
		if string(word[0]) > tree.Data {
			if tree.Right == nil {
				tree.Right = &types.Node{}
			}
			if tree.Right.Data == "" {
				insertIntoEmptySubtree(tree.Right, word, id, weight)
				return
			} else {
				tree = tree.Right
			}
		} else if string(word[0]) < tree.Data {
			if tree.Left == nil {
				tree.Left = &types.Node{}
			}
			if tree.Left.Data == "" {
				insertIntoEmptySubtree(tree.Left, word, id, weight)
				return
			} else {
				tree = tree.Left
			}
		} else if string(word[0]) == tree.Data {
			word = word[1:]
			if len(word) > 0 {
				if tree.Mid == nil {
					tree.Mid = &types.Node{}
				}
				if tree.Mid.Data == "" {
					insertIntoEmptySubtree(tree.Mid, word, id, weight)
					return
				} else {
					tree = tree.Mid
				}
			} else {
				var frequency int
				var weights []float32
				if tree.Occurances[id].Index == "" {
					frequency = 1
				} else {
					frequency = tree.Occurances[id].Frequency + 1
					weights = tree.Occurances[id].Weights
				}
				weights = append(weights, weight)
				tree.End = true
				tree.Occurances[id] = types.NodeOccurance{Index: id, Weights: weights, Frequency: frequency}
			}
		}
	}
	return
}
