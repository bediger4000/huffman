package tree

import (
	"fmt"
	"huffman/heap"
)

func (n *Interior) Value() float64 {
	return n.Freq
}

func (n *Interior) IsNil() bool {
	if n == nil {
		return true
	}
	return false
}

func (n *Interior) String() string {
	return fmt.Sprintf(":%.02f", n.Freq)
}

func (n *Interior) GreaterThan(other heap.Node) bool {
	if n.Freq > other.Value() {
		return true
	}
	return false
}
