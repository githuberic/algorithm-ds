package v0

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	arr := []int{20, 7, 3, 10, 15, 25, 30, 17, 19}
	n := len(arr)

	h := New(arr, n)
	h.Init()

	fmt.Println(h.Heap.String())
}

func TestHeapSort(t *testing.T) {
	arr := []int{20, 7, 3, 10, 15, 25, 30, 17, 19}
	n := len(arr)

	sortedArr := sort(arr, n)
	fmt.Printf("Sorted:%v", sortedArr)
}
