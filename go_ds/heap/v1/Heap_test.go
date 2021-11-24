package v1

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	var h = NewHeap()
	h.Val = []int{20, 7, 3, 10, 15, 25, 30, 17, 19}

	h.Init()
	fmt.Printf("Init,%v\n",h.Val) // [3 7 20 10 15 25 30 17 19]

	sorted := h.Sort()
	fmt.Printf("Sorted,%v\n",sorted)

	/*
	h.Push(6)
	fmt.Printf("Pushed,%v\n",h.Val)
*/
	x, ok := h.Remove(5)
	fmt.Println(x, ok, h.Val) // 25 true [3 6 15 10 7 20 30 17 19]

	y, ok := h.Remove(1)
	fmt.Println(y, ok, h.Val) // 6 true [3 7 15 10 19 20 30 17]

	z := h.Pop()
	fmt.Println(z, h.Val) // 3 [7 10 15 17 19 20 30]
}
