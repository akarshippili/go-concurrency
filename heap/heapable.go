package heap

type Heapable[T any] interface {
	Comparable[T]
	GetMin() T
}
