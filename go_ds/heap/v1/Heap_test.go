package v1

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	var h = NewHeap()
	h.Val = []int{20, 7, 3, 10, 15, 25, 30, 17, 19}
	h.Init()
	fmt.Println(h) // [3 7 20 10 15 25 30 17 19]

	h.Push(6)
	fmt.Println(h) // [3 6 20 10 7 25 30 17 19 15]

	x, ok := h.Remove(5)
	fmt.Println(x, ok, h) // 25 true [3 6 15 10 7 20 30 17 19]

	y, ok := h.Remove(1)
	fmt.Println(y, ok, h) // 6 true [3 7 15 10 19 20 30 17]

	z := h.Pop()
	fmt.Println(z, h) // 3 [7 10 15 17 19 20 30]
}
