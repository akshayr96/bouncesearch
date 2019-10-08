package types

type NodeOccurance struct {
	Index     string
	Weight    float32
	Frequency int
}

type NodeOccurances map[string]NodeOccurance

// Tree Node Struct
type Node struct {
	Left       *Node
	Right      *Node
	Mid        *Node
	Data       string
	End        bool
	Occurances NodeOccurances
}
