package heap

import "fmt"

type Heap struct {
	Arr      []int
	Capacity int
	Count    int
}

//init heap
func NewHeap(capacity int) *Heap {
	heap := &Heap{}
	heap.Capacity = capacity
	heap.Arr = make([]int, capacity+1)
	heap.Count = 0
	return heap
}

// swap two elements
func (h *Heap) Swap(i int, j int) {
	h.Arr[i], h.Arr[j] = h.Arr[j], h.Arr[i]
}

func (h *Heap) String() string {
	return fmt.Sprintf("Arr:%+v, Capacity:%+v, Count:%+v", h.Arr, h.Capacity, h.Count)
}
