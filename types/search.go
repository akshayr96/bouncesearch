package types

// Search Match Struct
type StringMatch int

const (
	NoMatch       StringMatch = 1
	PartialMatch  StringMatch = 2
	CompleteMatch StringMatch = 3
)
