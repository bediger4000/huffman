// Package tree represents a Huffman-coding tree of symbols (Go runes)
// and their frequencies. You can print a dot-format representation of
// a tree on stdout.
package tree

// Node interface allows the tree code to not care if it
// deals with a leaf or an interior node.
type Node interface {
	Value() float64
	IsNil() bool
	String() string
	LeftChild() Node
	RightChild() Node
}

// Leaf node in huffman-encoding tree. No children,
// holds a symbol and its frequency
type Leaf struct {
	Freq float64
	Char rune
}

// Interior node in huffman-encoding tree. Has children,
// and frequency of subtrees combined, but no children
type Interior struct {
	Freq  float64
	Left  Node
	Right Node
}
