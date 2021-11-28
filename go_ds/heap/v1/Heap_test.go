package v1

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	arr := []int{20, 7, 3, 10, 15, 25, 30, 17, 19}
	var h = NewHeapV1(arr)

	h.Init()
	fmt.Println("Init", h.String()) // [3 7 20 10 15 25 30 17 19]

	h.Push(6)
	fmt.Println("Pushed ", h.String())

	x, ok := h.Remove(5)
	fmt.Println(x, ok, h.String()) // 25 true [3 6 15 10 7 20 30 17 19]

	y, ok := h.Remove(1)
	fmt.Println(y, ok, h.String()) // 6 true [3 7 15 10 19 20 30 17]

	z := h.Pop()
	fmt.Println(z, h.String()) // 3 [7 10 15 17 19 20 30]
}
