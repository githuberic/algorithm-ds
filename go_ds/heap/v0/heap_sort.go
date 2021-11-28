package v0

import (
	"algorithm-ds/go_ds/heap"
)

//build arr heap
func buidHeap(a []int, n int) *HeapV0 {
	heap := heap.NewHeap(n)
	heap.Arr = a

	heapV0 := NewHeapV0(*heap)

	//heapify from the last parent node
	for i := n / 2; i >= 1; i-- {
		heapV0.down(i)
	}
	return heapV0
}

//sort by ascend, arr index begin from 1, has n elements
func sort(a []int, n int) {
	heapV0 := buidHeap(a, n)

	k := n - 1
	for k >= 1 {
		heapV0.Swap(0, k)
		heapV0.down(k)
		k--
	}
}
