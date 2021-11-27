package v0

import "algorithm-ds/go_ds/heap"

type HeapV0 struct {
	heap.Heap
}

func NewHeapV0(base heap.Heap) *HeapV0 {
	return &HeapV0{base}
}

//top-max heap -> heapify from down to up
func (h *HeapV0) insert(data int) {
	//defensive
	if h.Count == h.Capacity {
		return
	}

	h.Count++
	h.Arr[h.Count] = data

	//compare with parent node
	i := h.Count
	parent := i / 2
	for parent > 0 && h.Arr[parent] < h.Arr[i] {
		h.Swap(parent, i)
		i = parent
		parent = i / 2
	}
}

//heapify
func (h *HeapV0) down(i int) {
	for {
		maxIndex := i
		if h.Arr[i] < h.Arr[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= h.Count && h.Arr[maxIndex] < h.Arr[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		h.Swap(i, maxIndex)
		i = maxIndex
	}
}

//heapfify from up to down
func (h *HeapV0) removeMax() {
	//defensive
	if h.Count == 0 {
		return
	}

	//swap max and last
	h.Swap(1, h.Count)
	h.Count--

	//heapify from up to down
	h.down(0)
}
