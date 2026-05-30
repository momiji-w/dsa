package heap

type Heap struct {
	nodes []int
}

func parent(idx int) int     { return (idx - 1) / 2 }
func leftChild(idx int) int  { return (idx * 2) + 1 }
func rightChild(idx int) int { return (idx * 2) + 2 }

func (h *Heap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := parent(idx)
	if h.nodes[idx] < h.nodes[p] {
		h.nodes[idx], h.nodes[p] = h.nodes[p], h.nodes[idx]
		h.heapifyUp(p)
	}
}

func (h *Heap) heapifyDown(idx int) {
	lIdx := leftChild(idx)
	rIdx := rightChild(idx)

	if idx >= h.Len() || lIdx >= h.Len() {
		return
	}

	minIdx := lIdx
	if rIdx < len(h.nodes) && h.nodes[lIdx] > h.nodes[rIdx] {
		minIdx = rIdx
	}

	if minIdx < h.Len() && h.nodes[minIdx] < h.nodes[idx] {
		h.nodes[minIdx], h.nodes[idx] = h.nodes[idx], h.nodes[minIdx]
		h.heapifyDown(minIdx)
	}
}

func (h *Heap) Len() int {
	return len(h.nodes)
}

func (h *Heap) Insert(v int) {
	h.nodes = append(h.nodes, v)
	h.heapifyUp(h.Len() - 1)
}

func (h *Heap) Delete() int {
	if h.Len() == 0 {
		return -1
	}

	root := h.nodes[0]
	h.nodes[0] = h.nodes[h.Len()-1]
	h.nodes = h.nodes[:h.Len()-1]
	h.heapifyDown(0)

	return root
}
