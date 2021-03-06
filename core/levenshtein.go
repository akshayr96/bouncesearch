package core

// Computes the Levenshtien distance between two strings
// Takes two words to which the Levenshtein distance has to be computed
// Returns the Levenshtein distance between the two string
func LevenshteinDistance(wordOne, wordTwo string) (distance int) {
	if wordOne == "" {
		return len(wordTwo)
	}

	if wordTwo == "" {
		return len(wordOne)
	}

	if wordOne[0] == wordTwo[0] {
		return LevenshteinDistance(wordOne[1:], wordTwo[1:])
	}

	diagonalCost := LevenshteinDistance(wordOne[1:], wordTwo[1:])
	horizontalCost := LevenshteinDistance(wordOne, wordTwo[1:])
	VerticalCost := LevenshteinDistance(wordOne[1:], wordTwo)

	if diagonalCost > horizontalCost {
		diagonalCost = horizontalCost
	}

	if diagonalCost > VerticalCost {
		diagonalCost = VerticalCost
	}

	return diagonalCost + 1
}
