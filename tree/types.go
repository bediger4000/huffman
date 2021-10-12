package tree

type Node interface {
	Value() float64
	IsNil() bool
	String() string
	LeftChild() Node
	RightChild() Node
}

type Leaf struct {
	Freq float64
	Char rune
	// Parent *Interior
}

type Interior struct {
	Freq  float64
	Left  Node
	Right Node
	// Parent *Interior
}
