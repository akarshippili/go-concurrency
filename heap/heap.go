package heap

import (
	"errors"
)

// min-heap
type Heap[T Heapable[T]] struct {
	size int
	Arr  []T
}

func GetHeap[T Heapable[T]]() *Heap[T] {
	heap := Heap[T]{
		size: 0,
		Arr:  make([]T, 0),
	}

	return &heap
}

// Heap Operatrions
func (heap *Heap[T]) Add(val T) {
	// add element tio heap

	val.SetIndex(heap.size)
	if heap.size < len(heap.Arr) {
		heap.Arr[heap.size] = val
	} else {
		heap.Arr = append(heap.Arr, val)
	}
	heap.size += 1

	// shift to restore heap
	heap.HepifyUp(heap.size - 1)
}

func (heap *Heap[T]) Peek() (*T, error) {
	if heap.IsEmpty() {
		return nil, errors.New("invalid operation: heap is empty")
	}

	return &(heap.Arr[0]), nil
}

func (heap *Heap[T]) Pop() (*T, error) {
	if heap.IsEmpty() {
		return nil, errors.New("invalid operation: heap is empty")
	}

	heap.Swap(0, heap.size-1)
	heap.size -= 1

	// shift to restore heap
	heap.HepifyDown(0)
	return &heap.Arr[heap.size], nil
}

func (heap *Heap[T]) DecreaseKey(index int) error {
	if index >= heap.size {
		return errors.New("index out of bound")
	}

	heap.Arr[index] = heap.Arr[index].GetMin()
	heap.HepifyUp(index)
	return nil
}

func (heap *Heap[T]) DecreaseKeyWith(index int, new T) error {
	if index >= heap.size {
		return errors.New("index out of bound")
	}

	if heap.Arr[index].Compare(new) == -1 {
		return errors.New("updating value should be less than the current")
	}

	heap.Arr[index] = new
	heap.HepifyUp(index)
	return nil
}

func (heap *Heap[T]) IncreaseKeyWith(index int, new T) error {
	if index >= heap.size {
		return errors.New("index out of bound")
	}

	// current val is already > new val
	if heap.Arr[index].Compare(new) == 1 {
		return errors.New("updating value should be greater than the current")
	}

	heap.Arr[index] = new
	heap.HepifyDown(index)
	return nil
}

func (heap *Heap[T]) Delete(index int) error {
	if index >= heap.size {
		return errors.New("index out of bound")
	}

	heap.DecreaseKey(index)
	heap.Pop()
	return nil
}

func (heap *Heap[T]) IsEmpty() bool {
	return heap.size == 0
}

func (heap *Heap[T]) GetNumOfElements() int {
	return heap.size
}

// Helper Functions
func (heap *Heap[T]) Swap(index1 int, index2 int) error {
	if index1 >= heap.size {
		return errors.New("index1 is out of bound")
	}

	if index2 >= heap.size {
		return errors.New("index1 is out of bound")
	}

	heap.Arr[index1], heap.Arr[index2] = heap.Arr[index2], heap.Arr[index1]
	heap.Arr[index1].SetIndex(index1)
	heap.Arr[index2].SetIndex(index2)
	return nil
}

// shift-up to restore heap
func (heap *Heap[T]) HepifyUp(index int) {
	parentIndex := GetParentIndex(index)
	if heap.Arr[index].Compare(heap.Arr[parentIndex]) >= 0 {
		return
	}

	// swap
	heap.Swap(index, parentIndex)
	heap.HepifyUp(parentIndex)
}

// shift-down to restore heap
func (heap *Heap[T]) HepifyDown(index int) {
	smallest := index
	leftChildIndex := GetLeftChildIndex(index)
	rightChildIndex := GetRightChildIndex(index)

	if leftChildIndex < heap.size && heap.Arr[smallest].Compare(heap.Arr[leftChildIndex]) == 1 {
		smallest = leftChildIndex
	}

	if rightChildIndex < heap.size && heap.Arr[smallest].Compare(heap.Arr[rightChildIndex]) == 1 {
		smallest = rightChildIndex
	}

	if smallest == index {
		// if already in correct position
		return
	}

	// swap parant with smallest child
	heap.Swap(index, smallest)
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
