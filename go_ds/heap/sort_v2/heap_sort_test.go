package sort_v2

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	s := []int{-1,9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	fmt.Println(s[1:])
	HeapSort(s)
	fmt.Println(s[1:])
}
