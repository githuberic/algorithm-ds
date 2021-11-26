package v2

type Heap struct {
	Val []int
}

// 本实例构建最小堆 /**
func NewHeap() *Heap {
	return &Heap{make([]int, 0, 32)}
}

func (h *Heap) swap(i, j int) {
	h.Val[i], h.Val[j] = h.Val[j], h.Val[i]
}

func (h *Heap) less(i, j int) bool {
	return h.Val[i] < h.Val[j]
}

func (h *Heap) down(i int) {
	//存储u
	tmp := i
	n := len(h.Val)
	left := 2*i + 1
	if left < n && h.less(left, tmp) {
		tmp = left
	}

	right := 2*i + 2
	if right < n && h.less(right, tmp) {
		tmp = right
	}

	if tmp != i {
		h.swap(i, tmp)
		h.down(tmp)
	}
}
