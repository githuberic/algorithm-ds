package v0

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	arr := []int{20, 7, 3, 10, 15, 25, 30, 17, 19}
	n := len(arr)

	heapv0 := buidHeap(arr,n)
	fmt.Println(heapv0.Heap.String())

	sort(arr,n)
	fmt.Println(heapv0.Heap.String())
}
