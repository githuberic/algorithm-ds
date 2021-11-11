package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{8, 10, 2, 4, 3, 5, 1, 6, 7}
	fmt.Println(arr)
	heapSort(arr, len(arr))
	fmt.Println(arr)
}
