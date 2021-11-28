package heap

import "fmt"

type Heap struct {
	Arr []int
}

//init heap
func NewHeap(arr []int) *Heap {
	heap := &Heap{}
	heap.Arr = arr
	return heap
}

// swap two elements
func (h *Heap) Swap(i int, j int) {
	h.Arr[i], h.Arr[j] = h.Arr[j], h.Arr[i]
}

func (h *Heap) Less(i, j int) bool {
	return h.Arr[i] < h.Arr[j]
}

func (h *Heap) String() string {
	return fmt.Sprintf("Arr:%+v", h.Arr)
}
