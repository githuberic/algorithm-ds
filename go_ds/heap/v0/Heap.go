package v0

import "algorithm-ds/go_ds/heap"

type HeapV0 struct {
	heap.Heap
}

func NewHeapV0(base heap.Heap) *HeapV0 {
	return &HeapV0{base}
}

func New(arr []int, capacity int) *HeapV0 {
	heap := heap.NewHeap(capacity)
	heap.Arr = arr

	return NewHeapV0(*heap)
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
/*
func (h *HeapV0) down(i int) {
	for {
		maxIndex := i
		if i*2 <= h.Count && h.Arr[i] < h.Arr[i*2] {
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
}*/


func (h *HeapV0) down(i int) {
	for {
		// Left
		left := 2*i + 1
		if left >= len(h.Arr) {
			break // i已经是叶子结点了
		}

		// 从left,right中选择min,
		j := left
		if r := left + 1; r < len(h.Arr) && h.Less(r, left) {
			j = r // 右孩子
		}

		if h.Less(i, j) {
			break // 如果父结点比孩子结点小，则不交换
		}

		h.Swap(i, j) // 交换父结点和子结点
		i = j        //继续向下比较
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

func (h *HeapV0) Init() {
	n := len(h.Arr)

	// i > n/2-1 的结点为叶子结点本身已经是堆了
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}
