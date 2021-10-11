package heap

// Insert puts a node in the heap at the right hand
// side of the bottom row, which keeps the right
// tree shape. siftUp moves the new node up the tree
// as far as it should go.
func (h Heap) Insert(n Node) Heap {
	h = append(h, n)
	h.siftUp(len(h) - 1)
	return h
}

// siftUp checks that item at index idx in the partially-ordered
// array is correctly position, changes positions if need be,
// then calls itself on that newly positioned item's index.
func (h Heap) siftUp(idx int) {
	if idx == 0 {
		return
	}

	parentIdx := (idx - 1) / 2
	if !h[idx].GreaterThan(h[parentIdx]) {
		h[idx], h[parentIdx] = h[parentIdx], h[idx]
		h.siftUp(parentIdx)
	}
}
