package heap

type Heap struct {
	nodes []int
}

func (h *Heap) Len() int
func (h *Heap) Insert(v int) int
func (h *Heap) Delete() int
