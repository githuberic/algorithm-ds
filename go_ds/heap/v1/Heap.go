package v1

type Heap struct {
	Val []int
}

func NewHeap() *Heap {
	return &Heap{make([]int, 0, 32)}
}

func (h *Heap) swap(i, j int) {
	h.Val[i], h.Val[j] = h.Val[j], h.Val[i]
}

func (this *Heap) less(i, j int) bool {
	return this.Val[i] < this.Val[j]
}

func (this *Heap) up(i int) {
	for {
		// 父亲结点
		parent := (i - 1) / 2
		if i == parent || this.less(parent, i) {
			break
		}

		this.swap(parent, i)
		i = parent
	}
}

// 注意go中所有参数转递都是值传递
// 所以要让h的变化在函数外也起作用，此处得传指针
func (h *Heap) Push(x int) {
	h.Val = append(h.Val, x)
	h.up(len(h.Val) - 1)
}

func (h *Heap) Down(i int) {
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
		i = j        //继续向下比较
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
		h.Down(i)
	} else { // 当前元素小于父结点，向上筛选
		h.up(i)
	}
	return x, true
}

// 弹出堆顶的元素，并返回其值
func (h *Heap) Pop() int {
	n := len(h.Val) - 1
	h.swap(0, n)
	x := (h.Val)[n]
	h.Val = (h.Val)[0:n]
	h.Down(0)
	return x
}

func (h *Heap) Init() {
	n := len(h.Val)
	// i > n/2-1 的结点为叶子结点本身已经是堆了
	for i := n/2 - 1; i >= 0; i-- {
		h.Down(i)
	}
}
