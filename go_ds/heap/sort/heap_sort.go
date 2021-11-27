package sort

// 交换数组位置
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 调整堆,第i个结点
func heapify(arr []int, n, i int) {
	if i >= n {
		return
	}

	// left 左子结点,right右左子结点
	left := 2*i + 1
	right := 2*i + 2
	max := i

	if left < n && arr[left] > arr[max] {
		max = left
	}
	if right < n && arr[right] > arr[max] {
		max = right
	}

	if max != i {
		swap(arr, max, i)
		heapify(arr, n, max)
	}
}

// 创建堆(构建最大堆)
func buildHeap(arr []int, n int) {
	lastNode := n - 1            // 最后一个结点
	parent := (lastNode - 1) / 2 // 最后一个结点的父节点

	// 往最后一个父节点开始，不断往前创建结点
	for i := parent; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func heapSort(arr []int, n int) {
	buildHeap(arr, n)

	// 重复：第一个结点和最后一个结点交换位置，然后重新调整堆排序，i--去掉最后一个结点
	for i := n - 1; i >= 0; i-- {
		swap(arr, i, 0)
		heapify(arr, i, 0)
	}
}
