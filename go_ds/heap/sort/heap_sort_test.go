package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{8, 10, 2, 4, 3, 5, 1, 6, 7}
	fmt.Printf("Origin array:%v,\n", arr)
	buildHeap(arr, len(arr))
	fmt.Printf("Build heap:%v,\n", arr)
	heapSort(arr, len(arr))
	fmt.Printf("Heap sort:%v,\n", arr)
}
