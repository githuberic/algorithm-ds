package heap

type Heap struct {
	arr []int
	n   int
	count int
}

//init heap
func NewHeap(capacity int) *Heap {
	heap := &Heap{}
	heap.n = capacity
	heap.arr = make([]int, capacity+1)
	heap.count = 0
	return heap
}

//top-max heap -> heapify from down to up
func (heap *Heap) insert(data int) {
	//defensive
	if heap.count == heap.n {
		return
	}

	heap.count++
	heap.arr[heap.count] = data

	//compare with parent node
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.arr[parent] < heap.arr[i] {
		swap(heap.arr, parent, i)
		i = parent
		parent = i / 2
	}
}

//heapfify from up to down
func (heap *Heap) removeMax() {
	//defensive
	if heap.count == 0 {
		return
	}

	//swap max and last
	swap(heap.arr, 1, heap.count)
	heap.count--

	//heapify from up to down
	heapifyUpToDown(heap.arr, heap.count, 0)
}

//heapify
func heapifyUpToDown(a []int, count int, i int) {
	for i := 1; i <= count/2; {
		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		swap(a, i, maxIndex)
		i = maxIndex
	}
}

//swap two elements
func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}
