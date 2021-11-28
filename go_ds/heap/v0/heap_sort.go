package v0

import (
	"algorithm-ds/go_ds/heap"
)

// sort by ascend, arr index begin from 1, has n elements
func sort(a []int, n int) []int {
	heap := heap.NewHeap(n)
	heap.Arr = a

	heapV0 := NewHeapV0(*heap)
	//heapV0.Init()

	for i := n - 1; i >= 0; i-- {
		heapV0.Swap(i, 0)
		heapV0.down(i)
	}
	return heapV0.Arr
}
