package tree

import (
	"fmt"
	"huffman/heap"
)

func (l *Leaf) Value() float64 {
	return l.Freq
}

func (l *Leaf) IsNil() bool {
	if l == nil {
		return true
	}
	return false
}

func (l *Leaf) String() string {
	return fmt.Sprintf("%c:%.02f", l.Char, l.Freq)
}

func (l *Leaf) GreaterThan(other heap.Node) bool {
	if l.Freq > other.Value() {
		return true
	}
	return false
}

func (l *Leaf) LeftChild() Node {
	return (*Leaf)(nil)
}

func (l *Leaf) RightChild() Node {
	return (*Leaf)(nil)
}
