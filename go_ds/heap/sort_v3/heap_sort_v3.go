package sort_v3

// 本例为最小堆
// 最大堆只需要修改less函数即可
type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) up(i int) {
	for {
		f := (i - 1) / 2 // 父亲结点
		if i == f || h.less(f, i) {
			break
		}
		h.swap(f, i)
		i = f
	}
}

func (h Heap) down(i int) {
	for {
		l := 2*i + 1 // 左孩子
		if l >= len(h) {
			break // i已经是叶子结点了
		}

		j := l
		if r := l + 1; r < len(h) && h.less(r, l) {
			j = r // 右孩子
		}
		if h.less(i, j) {
			break // 如果父结点比孩子结点小，则不交换
		}
		h.swap(i, j) // 交换父结点和子结点
		i = j        //继续向下比较
	}
}

// 注意go中所有参数转递都是值传递
// 所以要让h的变化在函数外也起作用，此处得传指针
func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}
