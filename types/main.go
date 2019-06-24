package types

// Tree Node Struct
type Node struct {
	Left  *Node
	Right *Node
	Mid   *Node
	Data  string
	End   bool
}

// Search Match Struct
type StringMatch int

const (
	NoMatch       StringMatch = 1
	PartialMatch  StringMatch = 2
	CompleteMatch StringMatch = 3
)
