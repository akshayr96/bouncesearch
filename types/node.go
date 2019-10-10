package types

type TreeNodeWeight struct {
	Weight    float32
	Frequency int
}

type TreeNodeWeights map[string]TreeNodeWeight

type NodeOccurance struct {
	Index           string
	TreeNodeWeights TreeNodeWeights
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
