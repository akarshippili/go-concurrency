package heap

import "errors"

// min-heap
type Heap[T Heapable[T]] struct {
	size int
	arr  []T
}

func GetHeap[T Heapable[T]]() *Heap[T] {
	heap := Heap[T]{
		size: 0,
		arr:  make([]T, 0),
	}

	return &heap
}

// Heap Operatrions
func (heap *Heap[T]) Add(val T) {
	// add element tio heap
	heap.size += 1
	heap.arr = append(heap.arr, val)

	// shift to restore heap
	heap.HepifyUp(heap.size - 1)
}

func (heap *Heap[T]) Peek() (*T, error) {
	if heap.IsEmpty() {
		return nil, errors.New("invalid operation: heap is empty")
	}

	return &(heap.arr[0]), nil
}

func (heap *Heap[T]) Pop() (*T, error) {
	if heap.IsEmpty() {
		return nil, errors.New("invalid operation: heap is empty")
	}

	heap.arr[0], heap.arr[heap.size-1] = heap.arr[heap.size-1], heap.arr[0]
	heap.size -= 1

	// shift to restore heap
	heap.HepifyDown(0)
	return &heap.arr[heap.size], nil
}

func (heap *Heap[T]) DecreaseKey(index int) error {
	if index >= heap.size {
		return errors.New("index out of bound")
	}

	heap.arr[index] = heap.arr[index].GetMin()
	heap.HepifyUp(index)
	heap.Pop()
	return nil
}

func (heap *Heap[T]) IsEmpty() bool {
	return heap.size == 0
}

// Helper Functions

// shift-up to restore heap
func (heap *Heap[T]) HepifyUp(index int) {
	parentIndex := GetParentIndex(index)
	if heap.arr[index].Compare(heap.arr[parentIndex]) >= 0 {
		return
	}

	// swap
	heap.arr[index], heap.arr[parentIndex] = heap.arr[parentIndex], heap.arr[index]
	heap.HepifyUp(parentIndex)
}

// shift-down to restore heap
func (heap *Heap[T]) HepifyDown(index int) {
	smallest := index
	leftChildIndex := GetLeftChildIndex(index)
	rightChildIndex := GetRightChildIndex(index)

	if leftChildIndex < heap.size && heap.arr[smallest].Compare(heap.arr[leftChildIndex]) == 1 {
		smallest = leftChildIndex
	}

	if rightChildIndex < heap.size && heap.arr[smallest].Compare(heap.arr[rightChildIndex]) == 1 {
		smallest = rightChildIndex
	}

	if smallest == index {
		// if already in correct position
		return
	}

	// swap parant with smallest child
	heap.arr[index], heap.arr[smallest] = heap.arr[smallest], heap.arr[index]
	heap.HepifyDown(smallest)
}

func GetParentIndex(index int) int {
	return (index - 1) / 2
}

func GetLeftChildIndex(index int) int {
	return 2*index + 1
}

func GetRightChildIndex(index int) int {
	return 2*index + 2
}
