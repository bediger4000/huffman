package heap

// Delete removes the minimum-value item from the heap
// and returns the remaining heap and that min value
// item.
func (h Heap) Delete() (Heap, Node) {
	n := h[0]
	h[0] = h[len(h)-1]
	h = h[:len(h)-1]
	h.siftDown(0)
	return h, n
}

// siftDown checks for appropriate position of item
// in partially-ordered array at index idx.
// It moves items not in the right place.
func (h Heap) siftDown(idx int) {
	if idx > len(h)-1 {
		return
	}

	left := 2*idx + 1
	if left > len(h)-1 {
		return
	}
	if h[idx].GreaterThan(h[left]) {
		h[idx], h[left] = h[left], h[idx]
		h.siftDown(left)
	}

	right := 2*idx + 2
	if right > len(h)-1 {
		return
	}
	if h[idx].GreaterThan(h[right]) {
		h[idx], h[right] = h[right], h[idx]
		h.siftDown(right)
	}
}
