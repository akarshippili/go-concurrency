package heap

type Heapable[T any] interface {
	Comparable[T]
	GetMin() T
	GetIndex() int
	SetIndex(index int)
}
