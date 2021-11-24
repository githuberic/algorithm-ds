package v1

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

func (h *Heap) up(i int) {
	for {
		// 父亲结点
		parent := (i - 1) / 2
		if i == parent || h.less(parent, i) {
			break
		}

		h.swap(parent, i)
		i = parent
	}
}

// 注意go中所有参数转递都是值传递
// 所以要让h的变化在函数外也起作用，此处得传指针
func (h *Heap) Push(x int) {
	h.Val = append(h.Val, x)
	h.up(len(h.Val) - 1)
}

func (h *Heap) downV2(i int) {
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
		h.downV2(tmp)
	}
}

func (h *Heap) down(i int) {
	for {
		// Left
		left := 2*i + 1
		if left >= len(h.Val) {
			break // i已经是叶子结点了
		}

		j := left
		if r := left + 1; r < len(h.Val) && h.less(r, left) {
			j = r // 右孩子
		}

		if h.less(i, j) {
			break // 如果父结点比孩子结点小，则不交换
		}

		h.swap(i, j) // 交换父结点和子结点

		i = j //继续向下比较
	}
}

// 删除堆中位置为i的元素
// 返回被删元素的值
func (h *Heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(h.Val)-1 {
		return 0, false
	}

	n := len(h.Val) - 1
	h.swap(i, n) // 用最后的元素值替换被删除元素
	// 删除最后的元素
	x := (h.Val)[n]
	h.Val = (h.Val)[0:n]

	// 如果当前元素大于父结点，向下筛选
	if (h.Val)[i] > (h.Val)[(i-1)/2] {
		h.down(i)
	} else {
		// 当前元素小于父结点，向上筛选
		h.up(i)
	}

	return x, true
}

// Pop 弹出堆顶的元素，并返回其值 /**
func (h *Heap) Pop() int {
	n := len(h.Val) - 1
	h.swap(0, n)
	x := (h.Val)[n]
	h.Val = (h.Val)[0:n]
	h.down(0)
	return x
}

func (h *Heap) Sort() []interface{} {
	var res []interface{}
	for len(h.Val) != 0 {
		res = append(res, h.Pop())
	}
	return res
}

func (h *Heap) Init() {
	n := len(h.Val)

	// i > n/2-1 的结点为叶子结点本身已经是堆了
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}
