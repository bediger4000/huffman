package heap

import (
	"fmt"
	"io"
	"os"
)

// Draw puts GraphViz notation for a drawing of the heap
// on stdout.
func Draw(h Heap) {
	fmt.Fprintf(os.Stdout, "digraph g {\n")
	DrawNode(os.Stdout, h, 0, "N")
	fmt.Fprintf(os.Stdout, "\n}\n")
}

// DrawNode puts GraphViz output for a single heap node
// on stdout. It doesn't output dot-shaped nil-node items
// for empty children - GraphViz doesn't layout binary trees
// with the heap property oddly.
func DrawNode(out io.Writer, h Heap, idx int, prefix string) {
	if idx > len(h)-1 || h[idx].IsNil() {
		return
	}

	fmt.Fprintf(out, "%s%d [label=\"%v\"];\n", prefix, idx, h[idx])

	left := 2*idx + 1
	DrawNode(out, h, left, prefix)
	if left < len(h) && !h[left].IsNil() {
		fmt.Fprintf(out, "%s%d -> %s%d;\n", prefix, idx, prefix, left)
	}
	right := 2*idx + 2
	DrawNode(out, h, right, prefix)
	if right < len(h) && !h[right].IsNil() {
		fmt.Fprintf(out, "%s%d -> %s%d;\n", prefix, idx, prefix, right)
	}

}
