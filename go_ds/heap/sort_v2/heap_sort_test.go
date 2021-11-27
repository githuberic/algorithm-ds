package sort_v2

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{-1,9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	fmt.Printf("Origin array:%v,\n", arr)
	HeapSort(arr)
	fmt.Printf("Heap sort:%v,\n", arr)
}
